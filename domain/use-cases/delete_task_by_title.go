package use_cases

import (
	log "api-bff-golang/infrastructure/logger"
	"api-bff-golang/interfaces/gateways"
)

type DeleteTaskByTitleUseCaseInterface interface {
	Process(title string) (int64, error)
}

type DeleteTaskByTitleUseCase struct {
	taskGateway gateways.TaskGatewayInterface
}

func NewDeleteTaskByTitleUseCase(taskGateway gateways.TaskGatewayInterface) *DeleteTaskByTitleUseCase {
	return &DeleteTaskByTitleUseCase{taskGateway}
}

func (d *DeleteTaskByTitleUseCase) Process(title string) (int64, error) {
	log.Info("[delete_task_by_title_use_case] init use case with title %s", title)
	r, e := d.taskGateway.DeleteByTitle(title)
	if e != nil {
		log.Error("[delete_task_by_title_use_case] error deleting task by title %s message %s", title, e.Error())
		return 0, e
	}
	log.Info("[delete_task_by_title_use_case] task deleted successfully with title %s quantity %d", title, r)
	return r, nil
}
