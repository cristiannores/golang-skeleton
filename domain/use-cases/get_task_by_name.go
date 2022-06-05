package use_cases

import (
	"api-bff-golang/infrastructure/database/mongo/drivers/models"
	"api-bff-golang/infrastructure/database/mongo/drivers/repository"
	log "api-bff-golang/infrastructure/logger"
)

type GetTaskByNameUseCaseInterface interface {
	Process(title string) (models.TaskMongoModel, error)
}

type GetTaskByNameUseCase struct {
	taskRepository repository.TaskMongoRepositoryInterface
}

func NewGetTaskByNameUseCase(taskRepository repository.TaskMongoRepositoryInterface) *GetTaskByNameUseCase {
	return &GetTaskByNameUseCase{taskRepository}
}

func (t *GetTaskByNameUseCase) Process(title string) (models.TaskMongoModel, error) {
	log.Info("[get_task_by_name] init use case with title %s", title)
	r, e := t.taskRepository.GetByTitle(title)
	if e != nil {
		log.Error("[get_task_by_name] error getting task by title %s error %s", title, e.Error())
		return models.TaskMongoModel{}, nil
	}
	log.Info("[get_task_by_name] task found successfully with title %s", title)
	return r, nil
}
