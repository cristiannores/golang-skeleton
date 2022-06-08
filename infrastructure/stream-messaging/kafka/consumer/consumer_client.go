package consumer

import (
	log "api-bff-golang/infrastructure/logger"
	"context"
	"github.com/segmentio/kafka-go"
)

type ConsumerClientInterface interface {
	Consumer(ctx context.Context) <-chan IncomingMessage
}

type ConsumerClient struct {
	r *kafka.Reader
}
type IncomingMessage struct {
	Message []byte
	Key     string
}

func New(address []string, topic string, group string) *ConsumerClient {
	log.Info("creating consumer instance addr: %s  topic: %s group: %s", address, topic, group)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: address,
		GroupID: group,
		Topic:   topic,
	})
	return &ConsumerClient{r}
}

func (c *ConsumerClient) Consumer(ctx context.Context) <-chan IncomingMessage {
	messages := make(chan IncomingMessage)
	go func() {
		defer close(messages)

		for {
			m, err := c.r.ReadMessage(ctx)
			if err != nil {
				return
			}

			var e IncomingMessage
			e.Message = m.Value
			e.Key = string(m.Key)

			messages <- e
		}
	}()

	return messages
}

func (c *ConsumerClient) close() error {
	return c.r.Close()
}
