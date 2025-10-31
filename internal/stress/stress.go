package stress

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Report struct {
	TargetURL    string
	Total        int
	StatusCounts map[int]int
	Errors       int
	Duration     time.Duration
}

func (r *Report) FormatStressReport() string {
	var sb strings.Builder
	sb.WriteString("=== Stress Test Report ===\n")
	sb.WriteString(fmt.Sprintf("Target URL: %s\n", r.TargetURL))
	sb.WriteString(fmt.Sprintf("Total time: %s\n", r.Duration))
	sb.WriteString(fmt.Sprintf("Total requests: %d\n", r.Total))
	if len(r.StatusCounts) > 0 {
		sb.WriteString("Status codes:\n")
		for code, count := range r.StatusCounts {
			sb.WriteString(fmt.Sprintf("  %d: %d\n", code, count))
		}
	}
	sb.WriteString(fmt.Sprintf("Errors (network/timeouts/etc): %d\n", r.Errors))
	return sb.String()
}

func Run(url string, requests int, concurrency int) (*Report, error) {
	if concurrency < 1 {
		concurrency = 1
	}
	if requests < 1 {
		requests = 1
	}
	if concurrency > requests {
		concurrency = requests
	}

	jobs := make(chan struct{}, requests)
	results := make(chan int, requests)
	errorsCh := make(chan error, requests)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	client := &http.Client{Timeout: 10 * time.Second}
	worker := createStressTestWorker(&wg, url, jobs, errorsCh, results, client)

	for range concurrency {
		go worker()
	}

	start := time.Now()
	for range requests {
		jobs <- struct{}{}
	}
	close(jobs)

	go waitAndCloseChannels(&wg, results, errorsCh)

	report := &Report{TargetURL: url, Total: requests, StatusCounts: make(map[int]int)}
	for code := range results {
		report.StatusCounts[code]++
	}
	for range errorsCh {
		report.Errors++
	}

	report.Duration = time.Since(start)
	return report, nil
}

func createStressTestWorker(wg *sync.WaitGroup, url string, jobs <-chan struct{}, errorsCh chan<- error, results chan<- int, client *http.Client) func() {
	return func() {
		defer wg.Done()
		for range jobs {
			resp, err := client.Get(url)
			if err != nil {
				errorsCh <- err
				continue
			}
			code := resp.StatusCode
			resp.Body.Close()
			results <- code
		}
	}
}

func waitAndCloseChannels(wg *sync.WaitGroup, results chan int, errorsCh chan error) {
	wg.Wait()
	close(results)
	close(errorsCh)
}
