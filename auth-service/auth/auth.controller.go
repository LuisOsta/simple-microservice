package auth

import (
	"context"

	"github.com/auth-service/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "auth"

func getCollection() *mongo.Collection {
	database := db.GetDatabase()
	return database.Collection(COLLECTION_NAME)
}

type HandleLoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(c *gin.Context) {
	// gets the username and password from the request
	// get the collection instance
	// atempt to search for a document in the collection with a matching username and password
	// if the document is found, return the document
	// if the document is not found, return an error
	var requestBody HandleLoginRequestBody
	c.BindJSON(&requestBody)
	println(requestBody.Username, requestBody.Password)
	collection := getCollection()

	var credentials bson.D
	error := collection.FindOne(context.TODO(), bson.D{}).Decode(&credentials)

	if error != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, credentials)

}
