package config

import "os"

type Configuration struct {
	MONGO_URI             string
	JWT_SECRET            string
	USER_SERVICE_ENDPOINT string
}

func GetConfiguration() Configuration {

	return Configuration{
		MONGO_URI:             os.Getenv("MONGO_URI"),
		JWT_SECRET:            os.Getenv("JWT_SECRET"),
		USER_SERVICE_ENDPOINT: os.Getenv("USER_SERVICE_ENDPOINT"),
	}

}
