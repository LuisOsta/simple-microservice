package config

import "os"

type Configuration struct {
	MONGO_URI   string
	ROUTER_HOST string
}

func GetConfiguration() Configuration {

	return Configuration{
		MONGO_URI:   os.Getenv("MONGO_URI"),
		ROUTER_HOST: os.Getenv("ROUTER_HOST"),
	}

}
