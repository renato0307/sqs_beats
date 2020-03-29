package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/publisher"
)

type mockBatch struct {
	publisher.Batch
}

func (b mockBatch) Events() []publisher.Event {
	return make([]publisher.Event, 1)
}

func (b mockBatch) ACK() {
}

func (b mockBatch) RetryEvents(events []publisher.Event) {
}

type mockSQS struct {
	sqsiface.SQSAPI
}

func (svc mockSQS) SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	entries := make([]*sqs.SendMessageBatchResultEntry, 1)

	output := sqs.SendMessageBatchOutput{
		Successful: entries,
	}

	return &output, nil
}

type mockObserver struct {
	outputs.Observer
}

func (o mockObserver) NewBatch(numberOfEvents int) {
}
