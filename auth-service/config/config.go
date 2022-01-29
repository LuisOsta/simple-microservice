package config

import (
	"os"
	"strings"
)

type Configuration struct {
	MONGO_URI             string
	JWT_SECRET            string
	USER_SERVICE_ENDPOINT string
}

func GetConfiguration() Configuration {

	if strings.HasSuffix(os.Args[0], ".test") {
		return Configuration{
			MONGO_URI:             "mongodb://localhost:27017",
			JWT_SECRET:            "secret",
			USER_SERVICE_ENDPOINT: "http://localhost:8080",
		}
	} else {
		return Configuration{
			MONGO_URI:             os.Getenv("MONGO_URI"),
			JWT_SECRET:            os.Getenv("JWT_SECRET"),
			USER_SERVICE_ENDPOINT: os.Getenv("USER_SERVICE_ENDPOINT"),
		}
	}

}
