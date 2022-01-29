package config

import "os"

type Configuration struct {
	MONGO_URI  string
	JWT_SECRET string
}

func GetConfiguration() Configuration {

	return Configuration{
		MONGO_URI:  os.Getenv("MONGO_URI"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

}
