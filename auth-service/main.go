package main

import (
	"github.com/auth-service/db"
	"github.com/auth-service/router"
)

func main() {
	db.InitializeDatabaseConnection()
	router := router.ConfigureRouter()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	defer db.DisconnectClient()
}
