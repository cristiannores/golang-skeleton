package use_cases

import (
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/stream-messaging/kafka/producer"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type GetTaskAndSendUseCaseInterface interface {
	Process(title string) error
}

type GetTaskAndSendUseCase struct {
	producer              producer.ProducerInterface
	getTaskByTitleUseCase GetTaskByNameUseCaseInterface
}

func NewGetTaskAndSendUseCase(producer producer.ProducerInterface, getTaskByTitleUseCase GetTaskByNameUseCaseInterface) *GetTaskAndSendUseCase {
	return &GetTaskAndSendUseCase{producer, getTaskByTitleUseCase}
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

	e = g.producer.Produce(context.Background(), m, []byte("mi key"))

	if e != nil {
		log.Error("[get_task_and_send_use_case] error sending message to kafka %s error %s", m, e.Error())
		return e
	}

	return nil

}
