package auth

import (
	"context"

	"github.com/auth-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection() *mongo.Collection {
	database := db.GetDatabase()
	return database.Collection(COLLECTION_NAME)
}

func getUserByCredentials(username string, password string) (bson.D, error) {

	collection := getCollection()

	var credentials bson.D
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: username}, primitive.E{Key: "password", Value: password}}).Decode(&credentials)

	if err != nil {
		return nil, err
	}

	return credentials, nil
}
