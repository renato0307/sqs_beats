package sqs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	var batch = mockBatch{}
	var svc = mockSQS{}
	var observer = mockObserver{}

	c := client{
		beatName: "beat_name",
		codec:    newCodec("7.6"),
		index:    "index",
		observer: observer,
		queueURL: "queue_url",
		svc:      svc,
	}

	err := c.Publish(batch)

	assert.Nil(t, err)
}
