package main

import (
	"os"

	"github.com/elastic/beats/filebeat/cmd"
	inputs "github.com/elastic/beats/filebeat/input/default-inputs"

	_ "github.com/renato0307/sqs_beats/outputs/sqs"
)

func main() {
	if err := cmd.Filebeat(inputs.Init).Execute(); err != nil {
		os.Exit(1)
	}
}