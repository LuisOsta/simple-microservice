package config

import (
	"os"
	"strings"
)

type Configuration struct {
	MONGO_URI   string
	ROUTER_HOST string
	DB_NAME     string
}

func GetConfiguration() Configuration {

	// Best effort attempt to determine whether the execution environment is during testing.
	if strings.HasSuffix(os.Args[0], ".test") {
		return Configuration{
			MONGO_URI:   "mongodb://localhost:27017",
			ROUTER_HOST: "localhost:5000",
			DB_NAME:     "profile-service-test",
		}
	} else {
		return Configuration{
			MONGO_URI:   os.Getenv("MONGO_URI"),
			ROUTER_HOST: os.Getenv("ROUTER_HOST"),
			DB_NAME:     os.Getenv("DB_NAME"),
		}
	}

}
