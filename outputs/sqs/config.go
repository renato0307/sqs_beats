package sqs

import (
	"errors"
	"time"

	"github.com/elastic/beats/libbeat/common"
)

type sqsConfig struct {
	Region     string        `config:"region"`
	QueueURL   string        `config:"queue_url"`
	BatchSize  int           `config:"batch_size"`
	MaxRetries int           `config:"max_retries"`
	Timeout    time.Duration `config:"timeout"`
	Backoff    backoff       `config:"backoff"`
}

type backoff struct {
	Init time.Duration
	Max  time.Duration
}

const (
	defaultBatchSize = 1
	maxBatchSize     = 10
)

var (
	defaultConfig = sqsConfig{
		BatchSize:  defaultBatchSize,
		Timeout:    90 * time.Second,
		MaxRetries: 3,
		Backoff: backoff{
			Init: 1 * time.Second,
			Max:  60 * time.Second,
		},
	}
)

func (c *sqsConfig) Validate() error {
	if c.Region == "" {
		return errors.New("region is not defined")
	}

	if c.QueueURL == "" {
		return errors.New("queue_url is not defined")
	}

	if c.BatchSize > maxBatchSize || c.BatchSize < 1 {
		return errors.New("invalid batch size")
	}

	return nil
}

func readConfig(cfg *common.Config) (*sqsConfig, error) {
	c := defaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
