package use_cases

import (
	log "api-bff-golang/infrastructure/logger"
	kafkastreamer "api-bff-golang/infrastructure/stream-messaging/kafka"
	"api-bff-golang/shared/utils/config"
	"encoding/json"
	"fmt"
	"time"
)

type GetTaskAndSendUseCaseInterface interface {
	Process(title string) error
}

type GetTaskAndSendUseCase struct {
	kafka                 *kafkastreamer.KafkaStream
	getTaskByTitleUseCase GetTaskByNameUseCaseInterface
}

func NewGetTaskAndSendUseCase(kafka *kafkastreamer.KafkaStream, getTaskByTitleUseCase GetTaskByNameUseCaseInterface) *GetTaskAndSendUseCase {
	return &GetTaskAndSendUseCase{kafka, getTaskByTitleUseCase}
}

func (g *GetTaskAndSendUseCase) Process(title string) error {

	log.Info("[get_task_and_send_use_case] init use case with title %s", title)
	t, e := g.getTaskByTitleUseCase.Process(title)
	if e != nil {
		log.Error("[get_task_and_send_use_case] error getting task from database with title %s error %s", title, e)
		return e
	}

	currentTime := time.Now()

	t.Title = t.Title + " kafka"
	t.Tags = append(t.Tags, fmt.Sprintf("Send to kafka : %s", currentTime.Format("2006.01.02 15:04:05")))

	m, _ := json.Marshal(t)

	e = g.kafka.ProduceMessage(kafkastreamer.ProducerParams{
		Topic:   config.GetString("kafka.taskProducer.topic"),
		Message: string(m),
		Key:     "myKey",
	})

	if e != nil {
		log.Error("[get_task_and_send_use_case] error sending message to kafka %s error %s", m, e.Error())
		return e
	}

	return nil

}
