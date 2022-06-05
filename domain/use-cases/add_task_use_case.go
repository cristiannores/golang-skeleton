package use_cases

import (
	"api-bff-golang/domain/entities"
	"api-bff-golang/infrastructure/database/mongo/drivers/models"
	"api-bff-golang/infrastructure/database/mongo/drivers/repository"
	log "api-bff-golang/infrastructure/logger"
	"encoding/json"
)

type AddTaskUseCaseInterface interface {
	Process(t *entities.TaskEntity) (models.TaskMongoModel, error)
}
type AddTaskUseCase struct {
	taskRepository repository.TaskMongoRepositoryInterface
}

func NewAddTaskUseCase(taskRepository repository.TaskMongoRepositoryInterface) *AddTaskUseCase {
	return &AddTaskUseCase{taskRepository}
}

func (a *AddTaskUseCase) Process(t *entities.TaskEntity) (models.TaskMongoModel, error) {
	log.Info("[add_task_use_case] init use case with entity %v", t)
	m, _ := json.Marshal(t)
	var tm models.TaskMongoModel
	json.Unmarshal(m, &tm)
	log.Info("[add_task_use_case] calling to repository %v", t)
	id, e := a.taskRepository.Insert(tm)
	if e != nil {
		log.Error("[add_task_use_case] error trying insert to the collection %s", e.Error())
		return models.TaskMongoModel{}, nil
	}
	tm.ID = id
	log.Info("[add_task_use_case] task added successfully %v into database", tm)
	return tm, nil
}
