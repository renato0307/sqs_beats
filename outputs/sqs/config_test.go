package sqs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		Name  string
		Input sqsConfig
		Valid bool
	}{
		{"No config", sqsConfig{}, false},
		{"Just region", sqsConfig{Region: "eu-west-1"}, false},
		{"Just queue_name", sqsConfig{QueueName: "test_queue"}, false},
		{"Just region and queue_name", sqsConfig{QueueName: "test_queue", Region: "eu-west-1"}, false},
		{"With all required", sqsConfig{QueueName: "test_queue", Region: "eu-west-1", BatchSize: 1}, true},
	}

	for _, test := range tests {
		assert.Equal(t, test.Valid, test.Input.Validate() == nil, test.Name)
	}
}
