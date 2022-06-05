package use_cases

import (
	"api-bff-golang/infraestructure/database/mongo/drivers/repository"
	log "api-bff-golang/infraestructure/logger"
	"errors"
)

type DeleteTaskByTitleUseCaseInterface interface {
	Process(title string) (int64, error)
}

type DeleteTaskByTitleUseCase struct {
	taskRepository repository.TaskMongoRepositoryInterface
}

func NewDeleteTaskByTitleUseCase(taskRepository repository.TaskMongoRepositoryInterface) *DeleteTaskByTitleUseCase {
	return &DeleteTaskByTitleUseCase{taskRepository}
}

func (d *DeleteTaskByTitleUseCase) Process(title string) (int64, error) {
	log.Info("[delete_task_by_title_usecase] init use case with title %s", title)

	r, e := d.taskRepository.DeleteByTitle(title)
	if e != nil {
		log.Error("[delete_task_by_title_use_case] error deleting task by title %s message %s", title, e.Error())
		return 0, e
	}
	if r == 0 {
		log.Error("[delete_task_by_title_use_case] nothing to delete %s", title)
		return 0, errors.New("nothing to delete")
	}
	log.Info("[delete_task_by_title_use_case] task deleted successfully with title %s quantity %d", title, r)
	return r, nil
}
