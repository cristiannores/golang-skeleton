package infraestructure

import (
	use_cases "api-bff-golang/domain/use-cases"
	mongo "api-bff-golang/infraestructure/database/mongo/client"
	"api-bff-golang/infraestructure/database/mongo/drivers/repository"
	log "api-bff-golang/infraestructure/logger"
	"api-bff-golang/infraestructure/web"
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
	mongoClient, taskRepository := LoadDatabase()

	// add task functionality
	useCaseAddTask := use_cases.NewAddTaskUseCase(taskRepository)
	ctrlAddTask := controllers2.NewAddTaskController(useCaseAddTask)
	inputAddTask := inputs.NewAddTaskInput(ctrlAddTask)
	// get task functionality
	useCaseGetTask := use_cases.NewGetTaskByNameUseCase(taskRepository)
	ctrlGetTask := controllers2.NewGetTaskByNameController(useCaseGetTask)
	inputGetTask := inputs.NewGetTaskByTitleInput(ctrlGetTask)
	// find all task functionality
	useCaseFindAllTask := use_cases.NewFindAllTaskUseCase(taskRepository)
	ctrlFindAllTask := controllers2.NewFindAllTaskController(useCaseFindAllTask)
	inputFindAllTask := inputs.NewFindAllTaskInput(ctrlFindAllTask)
	// delete task by title functionality
	useCaseDeleteTaskByTitle := use_cases.NewDeleteTaskByTitleUseCase(taskRepository)
	ctrlDeleteTaskByTitle := controllers2.NewDeleteTaskByTitleController(useCaseDeleteTaskByTitle)
	inputDeleteTaskByTitle := inputs.NewDeleteTaskByTitleInput(ctrlDeleteTaskByTitle)

	web.InitRoutes(inputAddTask, inputGetTask, inputFindAllTask, inputDeleteTaskByTitle)

	log.Info("Setup dependencies ready .")

	return mongoClient
}
