package kafkastreamer

import (
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/shared/utils/config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestNewKafkaStream(t *testing.T) {
	//given
	getLocalConfig()

	kafka := testcontainers.NewLocalDockerCompose(
		[]string{"./../../../docker/kafka/test-docker-compose.yml"},
		strings.ToLower(uuid.New().String()),
	)
	fmt.Println(kafka)
	kafka.WithCommand([]string{"up", "-d"}).Invoke()

	time.Sleep(15 * time.Second)

	kafkaClient := NewKafkaStream()
	kafkaClient.ProduceMessage(ProducerParams{
		Topic:   "my-topic",
		Message: "test-message",
		Key:     "test-key",
	})
	defer destroyKafka(kafka)

	c := make(chan []byte)

	go kafkaClient.AddConsumer(
		ConsumerParams{
			Topic:    "my-topic",
			Consumer: "my-consumer",
			Callback: func(value []byte) error {
				fmt.Println("Capturing varibale")
				c <- value
				fmt.Println("Variable captured")
				return nil
			},
		},
	)
	x := <-c
	fmt.Println("processing assert")
	require.Equal(t, string(x), "test-message")

}
func destroyKafka(compose *testcontainers.LocalDockerCompose) {
	compose.Down()
	time.Sleep(1 * time.Second)
}

var validate *validator.Validate

func getLocalConfig() error {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./../../../app/shared/utils/config/")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	validateConfig()
	return nil
}

func validateConfig() error {

	file, _ := ioutil.ReadFile("./../../../app/shared/utils/config/config.json")
	config := config.Config{}

	_ = json.Unmarshal([]byte(file), &config)
	validate = validator.New()
	e := validate.Struct(&config)
	if e != nil {
		log.Fatal("Error in config file with message  %s", e.Error())
	}
	return nil

}
