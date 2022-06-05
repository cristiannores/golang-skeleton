package repository

import (
	"api-bff-golang/infraestructure/database/mongo/client"
	"api-bff-golang/infraestructure/database/mongo/drivers/models"
	log "api-bff-golang/infraestructure/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskMongoRepositoryInterface interface {
	Insert(task models.TaskMongoModel) (primitive.ObjectID, error)
	GetByTitle(title string) (models.TaskMongoModel, error)
	FindAll() ([]models.TaskMongoModel, error)
	DeleteByTitle(title string) (int64, error)
}
type TaskMongoRepository struct {
	client     *client.MongoClient
	collection *mongo.Collection
	ctx        context.Context
}

func NewTaskMongoRepository(client *client.MongoClient) *TaskMongoRepository {
	collection := client.GetCollection("tasks")
	return &TaskMongoRepository{
		client, collection, client.Ctx,
	}
}

func (t *TaskMongoRepository) Insert(task models.TaskMongoModel) (primitive.ObjectID, error) {
	i, e := t.collection.InsertOne(t.ctx, task)

	if e != nil {
		log.Error("[task_mongo_repository] error inserting %s", e.Error())
		return primitive.ObjectID{}, e
	}
	log.Info("[task_mongo_repository] task inserted successfully with id %s", i.InsertedID)
	return (i.InsertedID).(primitive.ObjectID), nil
}

func (t *TaskMongoRepository) GetByTitle(title string) (models.TaskMongoModel, error) {
	var taskModel models.TaskMongoModel
	e := t.collection.FindOne(t.ctx, bson.M{
		"title": title,
	}).Decode(&taskModel)
	if e != nil {
		log.Error("[task_mongo_repository] getting task by title %s", e.Error())
		return models.TaskMongoModel{}, e
	}
	log.Error("[task_mongo_repository] task found successfully by  title %s", title)
	return taskModel, nil
}

func (t *TaskMongoRepository) FindAll() ([]models.TaskMongoModel, error) {
	var tasks []models.TaskMongoModel
	opts := options.Find()
	opts.SetSort(bson.D{{"title", -1}})

	r, e := t.collection.Find(t.ctx, bson.M{}, opts)

	if e != nil {
		log.Error("[task_mongo_repository] error getting task document to struct %s")
		return []models.TaskMongoModel{}, e
	}
	if err := r.All(t.ctx, &tasks); err != nil {
		log.Error("[task_mongo_repository] error parsing database document to struct %s", err.Error())
		return []models.TaskMongoModel{}, e
	}
	log.Error("[task_mongo_repository] task documents found successfully")
	return tasks, nil
}

func (t *TaskMongoRepository) DeleteByTitle(title string) (int64, error) {
	d, e := t.collection.DeleteMany(t.ctx, bson.M{
		"title": title,
	})
	if e != nil {
		log.Error("error inserting %s", e.Error())
		return 0, e
	}
	return d.DeletedCount, e
}
