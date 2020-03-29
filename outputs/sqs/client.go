package sqs

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/codec"
	"github.com/elastic/beats/libbeat/outputs/codec/json"
	"github.com/elastic/beats/libbeat/publisher"
)

type client struct {
	svc      *sqs.SQS
	queueURL string
	beatName string
	index    string
	codec    codec.Codec
	timeout  time.Duration
	observer outputs.Observer
}

func newClient(sess *session.Session, config *sqsConfig, observer outputs.Observer, beat beat.Info) (*client, error) {
	client := &client{
		svc:      sqs.New(sess),
		queueURL: config.QueueURL,
		beatName: beat.Beat,
		index:    beat.IndexPrefix,
		codec: json.New(beat.Version, json.Config{
			Pretty:     false,
			EscapeHTML: false,
		}),
		timeout:  config.Timeout,
		observer: observer,
	}

	return client, nil
}

func (c client) String() string {
	return fmt.Sprintf("sqs(%s)", c.queueURL)
}

func (c *client) Close() error {
	return nil
}

func (c *client) Connect() error {
	return nil
}

func (c *client) Publish(batch publisher.Batch) error {

	// checks client and batch
	if c == nil {
		panic("no client")
	}
	if batch == nil {
		panic("no batch")
	}

	// gets the events to send
	events := batch.Events()
	c.observer.NewBatch(len(events))

	// converts events to sqs batch entries
	var entries []*sqs.SendMessageBatchRequestEntry
	entries = make([]*sqs.SendMessageBatchRequestEntry, len(events))
	for i, event := range events {
		serializedEvent, err := c.codec.Encode(c.index, &event.Content)
		if err != nil {
			return err
		}

		entries[i] = &sqs.SendMessageBatchRequestEntry{
			Id:          aws.String(string(i)),
			MessageBody: aws.String(string(base64Encode(serializedEvent))),
		}
	}

	// send the messages
	var input = &sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: &c.queueURL,
	}
	output, err := c.svc.SendMessageBatch(input)
	if err != nil {
		return err
	}

	// handles failed messages
	if len(output.Failed) > 0 {
		c.observer.Failed(len(output.Failed))
		var retryEvents []publisher.Event
		retryEvents = make([]publisher.Event, len(output.Failed))
		for i, failed := range output.Failed {
			n, _ := strconv.Atoi(*failed.Id)
			retryEvents[i] = events[n]
		}
		batch.RetryEvents(retryEvents)
		return nil
	}

	// if no message fails, ack the complete batch
	batch.ACK()
	return nil
}

func base64Encode(value []byte) []byte {
	encoded := make([]byte, base64.URLEncoding.EncodedLen(len(value)))
	base64.URLEncoding.Encode(encoded, value)
	return encoded
}

func base64Decode(value []byte) ([]byte, error) {
	decoded := make([]byte, base64.URLEncoding.DecodedLen(len(value)))
	b, err := base64.URLEncoding.Decode(decoded, value)
	if err != nil {
		return nil, err
	}
	return decoded[:b], nil
}
