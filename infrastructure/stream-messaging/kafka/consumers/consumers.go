package consumers

import (
	log "api-bff-golang/infrastructure/logger"
	kafkastreamer "api-bff-golang/infrastructure/stream-messaging/kafka"
	"api-bff-golang/interfaces/inputs"
	"api-bff-golang/shared/utils/config"
)

const APP = "golang-skeleton"

func InitConsumers(
	kafka *kafkastreamer.KafkaStream,
	addTask inputs.AddTaskInputInterface,

) {

	if config.GetBool("kafka.consumerExample.enable") {
		kafka.AddConsumer(kafkastreamer.ConsumerParams{
			Topic:    config.GetString("kafka.consumerExample.topic"),
			Consumer: config.GetString("kafka.consumerExample.prefix") + APP,
			Callback: addTask.FromKafka,
		})
	} else {
		log.Info("topic consumer disabled %s", config.GetString("kafka.consumeExample.topic"))
	}

}
