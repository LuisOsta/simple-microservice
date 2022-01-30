package config

import (
	"os"
	"strings"
)

type Configuration struct {
	MONGO_URI   string
	ROUTER_HOST string
}

func GetConfiguration() Configuration {

	// best effort attempt to determine whether the execution environment is during testing
	if strings.HasSuffix(os.Args[0], ".test") {
		return Configuration{
			MONGO_URI:   "mongodb://localhost:27017",
			ROUTER_HOST: "localhost:5000",
		}
	} else {
		return Configuration{
			MONGO_URI:   os.Getenv("MONGO_URI"),
			ROUTER_HOST: os.Getenv("ROUTER_HOST"),
		}
	}

}
