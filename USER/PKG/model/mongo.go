package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempy"`
	Name     string             `bson:"name, omitempty"`
	LastName string             `bson:"lastname, omitempty"`
}
