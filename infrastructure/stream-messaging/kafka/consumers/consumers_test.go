package consumers

import (
	"api-bff-golang/domain/entities"
	use_cases "api-bff-golang/domain/use-cases"
	"api-bff-golang/infrastructure/database/mongo/drivers/repository"
	"api-bff-golang/infrastructure/stream-messaging/kafka/consumer"
	controllers2 "api-bff-golang/interfaces/controllers"
	"api-bff-golang/interfaces/inputs"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/kafka"
	"github.com/orlangure/gnomock/preset/mongo"
	"time"

	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestFunctionalAddTaskByKafkaInput(t *testing.T) {
	gomega.RegisterFailHandler(Fail)
	RunSpecs(t, "functional kafka checkout order")
}

var _ = Describe("with containers & mocks gomock", Ordered, func() {

	var (
		useCaseAddTask     use_cases.AddTaskUseCaseInterface
		ctrl               *gomock.Controller
		ctrlAddTask        controllers2.AddTaskControllerInterface
		inputAddTask       inputs.AddTaskInputInterface
		taskRepositoryMock *repository.MockTaskMongoRepositoryInterface
		consumerClientMock *consumer.MockConsumerClientInterface
	)

	BeforeAll(func() {
		ctrl = gomock.NewController(GinkgoT())
		taskRepositoryMock = repository.NewMockTaskMongoRepositoryInterface(ctrl)
		consumerClientMock = consumer.NewMockConsumerClientInterface(ctrl)

	})
	When("input has correct parameters [with container]", func() {
		It("should return sucess  message", func() {

			// DATABASE
			pMongo := mongo.Preset(
				mongo.WithData("./testdata/"),
				mongo.WithUser("gnomock", "gnomick"),
			)
			ct, err := gnomock.Start(
				pMongo)

			defer func() { _ = gnomock.Stop(ct) }()

			if err != nil {
				panic(err)
			}

			addr := ct.DefaultAddress()
			uri := fmt.Sprintf("mongodb://%s:%s@%s", "gnomock", "gnomick", addr)
			clientOptions := mongooptions.Client().ApplyURI(uri)

			client, err := mongodb.NewClient(clientOptions)
			if err != nil {
				panic(err)
			}

			ctx := context.Background()

			err = client.Connect(ctx)
			if err != nil {
				panic(err)
			}

			// DATABASE
			taskRepository := repository.NewTaskMongoRepository(client)
			useCaseAddTask = use_cases.NewAddTaskUseCase(taskRepository)
			ctrlAddTask = controllers2.NewAddTaskController(useCaseAddTask)
			inputAddTask = inputs.NewAddTaskInput(ctrlAddTask)

			in := `{"title":"my-title","author":"my-author","tags":["my-tag1"]}`
			messages := []kafka.Message{
				{
					Topic: "events",
					Key:   "order",
					Value: in,
					Time:  time.Now().UnixNano(),
				},
			}

			p := kafka.Preset(
				kafka.WithTopics("topic-1", "topic-2"),
				kafka.WithMessages(messages...),
			)
			container, _ := gnomock.Start(
				p,
				gnomock.WithContainerName("kafka"),
			)

			_, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()
			c := consumer.New([]string{container.Address(kafka.BrokerPort)}, "events", "consumer11")

			result := AddTaskConsumer(inputAddTask, c)
			valueReceived := <-result

			Expect(valueReceived.result.Author).To(Equal("my-author"))
		})

	})

	When("input has correct parameters [with mock]", func() {
		It("should return success model", func() {
			lala := "my id"
			x, _ := primitive.ObjectIDFromHex(lala)
			taskRepositoryMock.EXPECT().Insert(gomock.Any()).Return(x, nil).AnyTimes()

			channel := make(chan consumer.IncomingMessage)
			o := entities.TaskEntity{
				Title:  "my-title",
				Author: "my-author",
				Tags:   []string{"my-tag1", "my-tag2"},
			}

			byteEntity, _ := json.Marshal(o)

			var message consumer.IncomingMessage
			message.Message = byteEntity
			message.Key = "any_key"
			go func() {
				channel <- message
			}()

			consumerClientMock.EXPECT().Consumer(gomock.Any()).Return(channel).AnyTimes()
			useCaseAddTask = use_cases.NewAddTaskUseCase(taskRepositoryMock)
			ctrlAddTask = controllers2.NewAddTaskController(useCaseAddTask)
			inputAddTask = inputs.NewAddTaskInput(ctrlAddTask)

			result := AddTaskConsumer(inputAddTask, consumerClientMock)
			valueReceived := <-result

			Expect(valueReceived.result.Author).To(ContainSubstring("my-author"))
			Expect(valueReceived.result.ID).To(Equal(x))
		})

	})

})
