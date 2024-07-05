package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"stress-test/application"
)

var url string
var requests uint64
var concurrency uint64
var rootCmd = &cobra.Command{
	Short: "Runs a http stress test",
	Long:  `A fast and flexible stress tester for http requests`,
	Run: func(cmd *cobra.Command, args []string) {
		stressService, err := application.NewStressTesterService(url, requests, concurrency)
		if err != nil {
			panic(err)
		}

		reportService := application.ReportPrinterService{}
		reportService.Print(stressService.Test())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "--url=http://google.com")
	rootCmd.PersistentFlags().Uint64Var(&requests, "requests", 0, "--requests=99")
	rootCmd.PersistentFlags().Uint64Var(&concurrency, "concurrency", 1, "--concurrency=5")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
