package sqs

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/outputs"
)

func init() {
	outputs.RegisterType("sqs", makeSqs)
}

func makeSqs(
	_ outputs.IndexManager,
	beat beat.Info,
	observer outputs.Observer,
	cfg *common.Config,
) (outputs.Group, error) {

	config, err := readConfig(cfg)
	if err != nil {
		return outputs.Fail(err)
	}

	client, err := newClient(config, observer, beat)
	if err != nil {
		return outputs.Fail(err)
	}

	retry := 0
	if config.MaxRetries < 0 {
		retry = -1
	}
	return outputs.Success(config.BatchSize, retry, client)
}
