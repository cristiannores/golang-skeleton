package consumers

import (
	"api-bff-golang/infrastructure/database/mongo/drivers/models"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/stream-messaging/kafka/consumer"
	"api-bff-golang/interfaces/inputs"
	"context"
)

const APP = "golang-skeleton"

type TaskResultKakfa struct {
	result models.TaskMongoModel
	e      error
}

func AddTaskConsumer(
	addTask inputs.AddTaskInputInterface,
	consumer consumer.ConsumerClientInterface,
) <-chan TaskResultKakfa {
	resultsTask := make(chan TaskResultKakfa)
	go func() {

		defer close(resultsTask)
		for m := range consumer.Consumer(context.Background()) {

			log.Info("consumer 1 reading message channel consumer message %s", string(m.Message))
			r, e := addTask.FromKafka(m.Message)

			var t TaskResultKakfa
			t.result = r
			t.e = e

			resultsTask <- t
		}
	}()
	return resultsTask

}
