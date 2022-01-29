package profile

import (
	"github.com/user-service/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "profile"

type ProfileDocument struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

func getCollection() *mongo.Collection {
	db := db.GetDatabase()
	return db.Collection(COLLECTION_NAME)
}

func createProfile(username string, password string) (ProfileDocument, error) {

	coll := getCollection()

	return ProfileDocument{}, nil
}
