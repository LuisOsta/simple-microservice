package db

import (
	"context"
	"log"
	"time"

	"github.com/user-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_TIMEOUT = 10 * time.Second

var database *mongo.Database = nil
var globalClient *mongo.Client = nil

func GetDatabase() *mongo.Database {

	if database != nil {
		log.Println("Database connection already established")
		return database
	}

	InitializeDatabaseConnection()

	return database
}

func InitializeDatabaseConnection() {
	log.Printf("Initializing database connection")

	ctx, cancel := context.WithTimeout(context.Background(), CONNECTION_TIMEOUT)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfiguration().MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finished establishing connection to the database")
	database = client.Database("test")
	globalClient = client
}

func DisconnectClient() {
	if globalClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := globalClient.Disconnect(ctx)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}

}
