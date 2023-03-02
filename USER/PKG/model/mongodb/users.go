package mongodb

import (
	"arshsuri96/cinema/USER/PKG/model"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	C *mongo.Collection
}

func (m UserModel) All() ([]model.User, error) {

	ctx := context.TODO()
	uu := []model.User{}
	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &uu); err != nil {
		log.Fatal(err)
	}
	return uu, err
}

func (m UserModel) FindByID(id string) (*model.User, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	var user = model.User{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&user)

	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}
