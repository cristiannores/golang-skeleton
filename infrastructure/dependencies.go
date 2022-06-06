package infrastructure

import (
	use_cases "api-bff-golang/domain/use-cases"
	mongo "api-bff-golang/infrastructure/database/mongo/client"
	"api-bff-golang/infrastructure/database/mongo/drivers/repository"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/infrastructure/metrics/prometheus"
	kafkastreamer "api-bff-golang/infrastructure/stream-messaging/kafka"
	"api-bff-golang/infrastructure/stream-messaging/kafka/consumers"
	"api-bff-golang/infrastructure/web"
	controllers2 "api-bff-golang/interfaces/controllers"
	"api-bff-golang/interfaces/inputs"
)

func LoadDatabase() (*mongo.MongoClient, repository.TaskMongoRepositoryInterface) {
	m := mongo.NewMongoClient()
	_, e := m.Connect()
	if e != nil {
		log.Fatal("error connecting to database")
	}
	tr := repository.NewTaskMongoRepository(m)
	return m, tr
}

func SetupDependencies() *mongo.MongoClient {
	log.Info("Setup dependencies ...")
	//modules
	mongoClient, taskRepository := LoadDatabase()
	kafka := kafkastreamer.NewKafkaStream()
	metric := prometheus.NewMetric()
	metric.InitMetrics()
	// add task functionality
	useCaseAddTask := use_cases.NewAddTaskUseCase(taskRepository)
	ctrlAddTask := controllers2.NewAddTaskController(useCaseAddTask)
	inputAddTask := inputs.NewAddTaskInput(ctrlAddTask)
	// get task functionality
	useCaseGetTask := use_cases.NewGetTaskByNameUseCase(taskRepository)
	ctrlGetTask := controllers2.NewGetTaskByNameController(useCaseGetTask, metric)
	inputGetTask := inputs.NewGetTaskByTitleInput(ctrlGetTask)
	// find all task functionality
	useCaseFindAllTask := use_cases.NewFindAllTaskUseCase(taskRepository)
	ctrlFindAllTask := controllers2.NewFindAllTaskController(useCaseFindAllTask)
	inputFindAllTask := inputs.NewFindAllTaskInput(ctrlFindAllTask)
	// delete task by title functionality
	useCaseDeleteTaskByTitle := use_cases.NewDeleteTaskByTitleUseCase(taskRepository)
	ctrlDeleteTaskByTitle := controllers2.NewDeleteTaskByTitleController(useCaseDeleteTaskByTitle)
	inputDeleteTaskByTitle := inputs.NewDeleteTaskByTitleInput(ctrlDeleteTaskByTitle)

	// get task and send functionality
	useCaseGetTaskAndSend := use_cases.NewGetTaskAndSendUseCase(kafka, useCaseGetTask)
	ctrlGetTaskAndSend := controllers2.NewGetTaskAndSendController(useCaseGetTaskAndSend)
	inputGetTaskAndSend := inputs.NewGetLastTaskAndSendInput(ctrlGetTaskAndSend)

	web.InitRoutes(inputAddTask, inputGetTask, inputFindAllTask, inputDeleteTaskByTitle, inputGetTaskAndSend)

	go consumers.InitConsumers(kafka, inputAddTask)
	log.Info("Setup dependencies ready .")

	return mongoClient
}
