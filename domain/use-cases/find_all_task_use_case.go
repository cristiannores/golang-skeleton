package use_cases

import (
	"api-bff-golang/infrastructure/database/mongo/drivers/models"
	"api-bff-golang/infrastructure/database/mongo/drivers/repository"
	log "api-bff-golang/infrastructure/logger"
)

type FindAllTaskUseCaseInterface interface {
	Process() ([]models.TaskMongoModel, error)
}

type FindAllTaskUseCase struct {
	taskRepository repository.TaskMongoRepositoryInterface
}

func NewFindAllTaskUseCase(taskRepository repository.TaskMongoRepositoryInterface) *FindAllTaskUseCase {
	return &FindAllTaskUseCase{taskRepository}
}

func (f *FindAllTaskUseCase) Process() ([]models.TaskMongoModel, error) {
	log.Info("[find_all_task_use_case] init use case")

	r, e := f.taskRepository.FindAll()
	if e != nil {
		log.Error("[find_all_task_use_case] error geting all task %s", e.Error())
		return []models.TaskMongoModel{}, e
	}
	log.Info("[find_all_task_use_case] tasks found successfully length  %d", len(r))
	return r, nil
}
