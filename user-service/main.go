package main

import "github.com/user-service/router"

func main() {
	router := router.ConfigureRouter()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
