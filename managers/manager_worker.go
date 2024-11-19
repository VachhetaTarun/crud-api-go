package repository

import (
	"context"
	models "crudecho/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkerRepository struct {
	client *mongo.Client
}

// Create a new worker
func (repo *WorkerRepository) CreateWorker(worker models.Worker) (*mongo.InsertOneResult, error) {
	collection := repo.client.Database("testdb").Collection("workers")
	return collection.InsertOne(context.Background(), worker)
}

// Get all workers
func (repo *WorkerRepository) GetAllWorkers() ([]models.Worker, error) {
	var workers []models.Worker
	collection := repo.client.Database("testdb").Collection("workers")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &workers); err != nil {
		return nil, err
	}
	return workers, nil
}

// Get worker by ID
func (repo *WorkerRepository) GetWorkerByID(id primitive.ObjectID) (models.Worker, error) {
	var worker models.Worker
	collection := repo.client.Database("testdb").Collection("workers")
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&worker)
	return worker, err
}

// Update worker by ID
func (repo *WorkerRepository) UpdateWorker(id primitive.ObjectID, update models.Worker) (*mongo.UpdateResult, error) {
	collection := repo.client.Database("testdb").Collection("workers")
	return collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": update})
}

// Delete worker by ID
func (repo *WorkerRepository) DeleteWorker(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := repo.client.Database("testdb").Collection("workers")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}
