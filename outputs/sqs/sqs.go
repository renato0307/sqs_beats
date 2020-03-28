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

	var group outputs.Group

	return group, nil
}
