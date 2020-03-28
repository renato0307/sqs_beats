package sqs

import (
	"errors"
	"time"
)

type sqsConfig struct {
	Region     string        `config:"region"`
	QueueName  string        `config:"queue_name"`
	BatchSize  int           `config:"batch_size"`
	MaxRetries int           `config:"max_retries"`
	Timeout    time.Duration `config:"timeout"`
	Backoff    backoff       `config:"backoff"`
}

type backoff struct {
	Init time.Duration
	Max  time.Duration
}

// TODO: review for SQS
const (
	defaultBatchSize = 50
	maxBatchSize     = 500
)

var (
	defaultConfig = sqsConfig{
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

	if c.QueueName == "" {
		return errors.New("queue_name is not defined")
	}

	if c.BatchSize > maxBatchSize || c.BatchSize < 1 {
		return errors.New("invalid batch size")
	}

	return nil
}
