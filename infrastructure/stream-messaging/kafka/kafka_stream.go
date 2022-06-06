package kafkastreamer

import (
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/shared/utils/config"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

type ConsumerParams struct {
	Topic    string
	Consumer string
	Callback func([]byte) error
}
type ProducerParams struct {
	Topic   string
	Message string
	Key     string
}
type KafkaStreamInterface interface {
	AddConsumer(params ConsumerParams) error
	ProduceMessage(params ProducerParams) error
	Close() error
}

// KafkaStream here is my method definition
type KafkaStream struct {
	consumers []*kafka.Reader
}

// NewKafkaStream This is my constructor ( convention New in constructor name )
func NewKafkaStream() *KafkaStream {
	return &KafkaStream{}
}
func (k *KafkaStream) AddConsumer(params ConsumerParams) error {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  config.GetArray("kafka.brokers"),
		GroupID:  params.Consumer,
		Topic:    params.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	log.Info("Kafka consumer added successfully Topic: %s | Consumer: %s", params.Topic, params.Consumer)
	k.consumers = append(k.consumers, r)
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		log.Info("[ReadMessage] message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		e := params.Callback(m.Value)
		if e != nil {
			// prevent throw error
			log.Error("[ReadMessage] error reading message from topic %v value %s with error", m.Topic, string(m.Value), e.Error())
		}
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader: %s", err.Error())
	}
	return nil
}
func (k *KafkaStream) ProduceMessage(params ProducerParams) error {
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", params.Topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader: %s", err.Error())
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(params.Message), Key: []byte(params.Key)},
	)
	if err != nil {
		log.Fatal("failed to write messages: %s", err.Error())
	}
	log.Info("[ProduceMessage] Message sent correctly topic:[%s] | Message:[%s]", params.Topic, params.Message)
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer: %s", err.Error())
	}
	return nil
}
func (k *KafkaStream) Close() error {
	for _, consumer := range k.consumers {

		consumer.Close()
		log.Info("Closed consumer for topic : %s", consumer.Config().Topic)
	}

	return nil
}
