package auth

import (
	"context"
	"log"

	"github.com/auth-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthDocument struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

func getCollection() *mongo.Collection {
	db := db.GetDatabase()
	return db.Collection(COLLECTION_NAME)
}

func getUserByCredentials(username string, password string) (AuthDocument, error) {

	coll := getCollection()

	var credentials AuthDocument
	err := coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: username}, primitive.E{Key: "password", Value: password}}).Decode(&credentials)

	if err != nil {
		log.Printf("Error finding user with credentials: %s", err.Error())
		return AuthDocument{}, err
	}

	return credentials, nil
}
