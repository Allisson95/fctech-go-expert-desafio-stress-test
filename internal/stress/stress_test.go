package stress

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRun_Simple200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	report, err := Run(srv.URL, 20, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.StatusCounts[200] != 20 {
		t.Fatalf("expected 20 success responses, got %d", report.StatusCounts[200])
	}
	if report.Total != 20 {
		t.Fatalf("expected total 20, got %d", report.Total)
	}
}

func TestRun_MixedStatusCodes(t *testing.T) {
	count := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		if count%3 == 0 {
			w.WriteHeader(404)
		} else if count%5 == 0 {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
		}
		w.Write([]byte("response"))
	}))
	defer srv.Close()

	report, err := Run(srv.URL, 15, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.Total != 15 {
		t.Fatalf("expected total 15, got %d", report.Total)
	}
	if report.StatusCounts[200] == 0 {
		t.Fatalf("expected some 200 responses")
	}
	if report.StatusCounts[404] == 0 {
		t.Fatalf("expected some 404 responses")
	}
}
