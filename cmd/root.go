package cmd

import (
	"os"

	"github.com/allisson95/fctech-go-expert-desafio-stress-test/internal/stress"
	"github.com/spf13/cobra"
)

var (
	urlFlag         string
	requestsFlag    int
	concurrencyFlag int
)

var rootCmd = &cobra.Command{
	Use:   "stressr",
	Short: "A small CLI to run HTTP stress tests",
	Long:  `stressr executes concurrent HTTP requests against a target URL and reports metrics.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if urlFlag == "" {
			cmd.PrintErr("Error: URL flag is required\n")
			return cmd.Help()
		}
		if requestsFlag <= 0 {
			requestsFlag = 1
		}
		if concurrencyFlag <= 0 {
			concurrencyFlag = 1
		}

		report, err := stress.Run(urlFlag, requestsFlag, concurrencyFlag)
		if err != nil {
			return err
		}

		cmd.Println(report.FormatStressReport())

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&urlFlag, "url", "", "URL of the service to be tested")
	rootCmd.PersistentFlags().IntVar(&requestsFlag, "requests", 1, "Total number of requests to perform")
	rootCmd.PersistentFlags().IntVar(&concurrencyFlag, "concurrency", 1, "Number of simultaneous calls")
}
