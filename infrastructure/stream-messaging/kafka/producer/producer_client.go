package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type ProducerInterface interface {
	Produce(ctx context.Context, key, value []byte) error
	Close() error
}
type Producer struct {
	w *kafka.Writer
}

func New(address []string, topic string) *Producer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(address...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{w}
}

func (p *Producer) Produce(ctx context.Context, key, value []byte) error {
	if err := p.w.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	}); err != nil {
		return err
	}

	return nil
}

func (p *Producer) Close() error {
	return p.w.Close()
}
