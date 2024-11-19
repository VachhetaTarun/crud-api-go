package services

import (
	"context"
	"crudecho/config"
	models "crudecho/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkerService struct {
	client *mongo.Client
}

func NewWorkerService(cfg *config.Config) *WorkerService {
	return &WorkerService{client: cfg.MongoClient}
}

func (s *WorkerService) CreateWorker(worker models.Worker) (*mongo.InsertOneResult, error) {
	collection := s.client.Database("testdb").Collection("workers")
	return collection.InsertOne(context.Background(), worker)
}

func (s *WorkerService) GetWorkerByID(id primitive.ObjectID) (models.Worker, error) {
	var worker models.Worker
	collection := s.client.Database("testdb").Collection("workers")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&worker)
	return worker, err
}

func (s *WorkerService) GetAllWorkers() ([]models.Worker, error) {
	var workers []models.Worker
	collection := s.client.Database("testdb").Collection("workers")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &workers); err != nil {
		return nil, err
	}
	return workers, nil
}

func (s *WorkerService) UpdateWorker(id primitive.ObjectID, update models.Worker) (*mongo.UpdateResult, error) {
	collection := s.client.Database("testdb").Collection("workers")
	return collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": update})
}

func (s *WorkerService) DeleteWorker(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := s.client.Database("testdb").Collection("workers")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}
