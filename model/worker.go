package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Worker struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Age   int                `bson:"age"`
}
