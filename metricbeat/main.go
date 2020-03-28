package main

import (
	"os"

	"github.com/elastic/beats/metricbeat/cmd"
	_ "github.com/renato0307/sqsbeatoutput/outputs/sqs"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
