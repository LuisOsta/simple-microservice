package proxy

import (
	"errors"
	"strings"

	"github.com/auth-service/config"
)

type service struct {
	Name     string
	Endpoint string
}

var services = [...]service{{Name: "user", Endpoint: config.GetConfiguration().USER_SERVICE_ENDPOINT}}

func getService(serviceName string) (service, error) {
	for _, service := range services {
		if service.Name == serviceName {
			return service, nil
		}
	}
	return service{}, errors.New("service not found")
}

func getServiceNameAndPath(path string) (string, string) {
	for _, service := range services {
		if strings.HasPrefix(path, "/"+service.Name) {
			return service.Name, path[len(service.Name)+1:]
		}
	}
	return "", path
}
