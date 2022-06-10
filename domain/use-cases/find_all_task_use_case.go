package use_cases

import (
	"api-bff-golang/domain/entities"
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/interfaces/gateways"
)

type FindAllTaskUseCaseInterface interface {
	Process() ([]entities.TaskEntity, error)
}

type FindAllTaskUseCase struct {
	taskGateway gateways.TaskGatewayInterface
}

func NewFindAllTaskUseCase(taskGateway gateways.TaskGatewayInterface) *FindAllTaskUseCase {
	return &FindAllTaskUseCase{taskGateway}
}

func (f *FindAllTaskUseCase) Process() ([]entities.TaskEntity, error) {
	log.Info("[find_all_task_use_case] init use case")

	r, e := f.taskGateway.FindAll()
	if e != nil {
		log.Error("[find_all_task_use_case] error getting all task %s", e.Error())
		return []entities.TaskEntity{}, e
	}
	log.Info("[find_all_task_use_case] tasks found successfully length  %d", len(r))
	return r, nil
}
