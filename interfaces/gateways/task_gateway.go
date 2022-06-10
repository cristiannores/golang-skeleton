package gateways

import (
	"api-bff-golang/domain/entities"
	"api-bff-golang/infrastructure/database/mongo/drivers/models"
	"api-bff-golang/infrastructure/database/mongo/drivers/repository"
	log "api-bff-golang/infrastructure/logger"
	"errors"
)

type TaskGatewayInterface interface {
	SaveTask(task *entities.TaskEntity) (entities.TaskEntity, error)
	DeleteByTitle(title string) (int64, error)
	FindAll() ([]entities.TaskEntity, error)
	GetByTitle(title string) (entities.TaskEntity, error)
}

type TaskGateway struct {
	taskRepository repository.TaskMongoRepositoryInterface
}

func NewTaskGateway(taskRepository repository.TaskMongoRepositoryInterface) *TaskGateway {
	return &TaskGateway{taskRepository}
}

func (t *TaskGateway) SaveTask(task *entities.TaskEntity) (entities.TaskEntity, error) {

	taskModel := generateTaskInitialModel(task)
	r, e := t.taskRepository.Insert(taskModel)
	if e != nil {
		log.Error("error adding task entity")
		return entities.TaskEntity{}, e
	}

	log.Info("task added successfully with id %v", r)
	return generateTaskEntity(taskModel), nil
}

func (t *TaskGateway) DeleteByTitle(title string) (int64, error) {

	r, e := t.taskRepository.DeleteByTitle(title)
	if e != nil {
		log.Error("error adding task entity")
		return 0, e
	}

	if r == 0 {
		return 0, errors.New("Error deleting ")
	}

	log.Info("task added successfully with id %v", r)
	return r, nil
}

func (t *TaskGateway) FindAll() ([]entities.TaskEntity, error) {

	tasks, e := t.taskRepository.FindAll()

	if e != nil {
		log.Error("Error finding tasks %s", e.Error())
		return []entities.TaskEntity{}, e
	}

	var taskEntities []entities.TaskEntity
	for _, task := range tasks {
		taskEntities = append(taskEntities, generateTaskEntity(task))
	}

	return taskEntities, nil
}

func (t *TaskGateway) GetByTitle(title string) (entities.TaskEntity, error) {

	r, e := t.taskRepository.GetByTitle(title)

	if e != nil {
		return entities.TaskEntity{}, e
	}
	return generateTaskEntity(r), nil
}
func generateTaskInitialModel(task *entities.TaskEntity) models.TaskMongoModel {

	return models.TaskMongoModel{
		Title:  task.Title,
		Author: task.Author,
		Tags:   task.Tags,
	}
}

func generateTaskEntity(taskModel models.TaskMongoModel) entities.TaskEntity {
	return entities.TaskEntity{
		Title:  taskModel.Title,
		Tags:   taskModel.Tags,
		Author: taskModel.Author,
	}
}
