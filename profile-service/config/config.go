package config

import "os"

type Configuration struct {
	MONGO_URI string
}

func GetConfiguration() Configuration {

	return Configuration{
		MONGO_URI: os.Getenv("MONGO_URI"),
	}

}
