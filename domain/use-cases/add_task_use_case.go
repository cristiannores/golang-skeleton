package use_cases

import (
	"api-bff-golang/domain/entities"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/interfaces/gateways"
	"api-bff-golang/shared/errors"
)

type AddTaskUseCaseInterface interface {
	Process(t *entities.TaskEntity) (entities.TaskEntity, error)
}
type AddTaskUseCase struct {
	taskGateway gateways.TaskGatewayInterface
}

func NewAddTaskUseCase(taskGateway gateways.TaskGatewayInterface) *AddTaskUseCase {
	return &AddTaskUseCase{taskGateway}
}

func (a *AddTaskUseCase) Process(t *entities.TaskEntity) (entities.TaskEntity, error) {
	log.Info("[add_task_use_case] init use case with entity %v", t)

	log.Info("[add_task_use_case] calling to repository %v", t)
	task, e := a.taskGateway.SaveTask(t)
	if e != nil {
		log.Error("[add_task_use_case] error trying to save task %s", e.Error())
		return entities.TaskEntity{}, errors.New([]string{t.Title}, e.Error(), errors.TASK_NOT_INSERTED)
	}

	log.Info("[add_task_use_case] task added successfully %v into database", t)
	return task, nil
}
