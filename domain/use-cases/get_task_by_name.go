package use_cases

import (
	"api-bff-golang/domain/entities"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/interfaces/gateways"
	"api-bff-golang/shared/errors"
)

type GetTaskByNameUseCaseInterface interface {
	Process(title string) (entities.TaskEntity, error)
}

type GetTaskByNameUseCase struct {
	taskGateway gateways.TaskGatewayInterface
}

func NewGetTaskByNameUseCase(taskGateway gateways.TaskGatewayInterface) *GetTaskByNameUseCase {
	return &GetTaskByNameUseCase{taskGateway}
}

func (t *GetTaskByNameUseCase) Process(title string) (entities.TaskEntity, error) {
	log.Info("[get_task_by_name] init use case with title %s", title)
	r, e := t.taskGateway.GetByTitle(title)
	if e != nil {
		log.Error("[get_task_by_name] error getting task by title %s error %s", title, e.Error())
		var s = []string{title}
		return entities.TaskEntity{}, errors.New(s, e.Error(), errors.TASK_NOT_FOUND)
	}
	log.Info("[get_task_by_name] task found successfully with title %s", title)
	return r, nil
}
