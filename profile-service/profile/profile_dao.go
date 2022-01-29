package profile

import (
	"context"

	"github.com/user-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "profile"

type ProfileDocument struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Address string             `bson:"address"`
	Phone   string             `bson:"phone"`
}

func getCollection() *mongo.Collection {
	db := db.GetDatabase()
	return db.Collection(COLLECTION_NAME)
}

func createProfile(address string, phone string) (ProfileDocument, error) {

	coll := getCollection()

	res, err := coll.InsertOne(context.TODO(), bson.D{{Key: "address", Value: address}, {Key: "phone", Value: phone}})

	if err != nil {
		return ProfileDocument{}, err
	}

	return ProfileDocument{ID: res.InsertedID.(primitive.ObjectID), Address: address, Phone: phone}, nil
}
