package proxy

import (
	"fmt"
	"log"
	"strings"

	"github.com/auth-service/config"
)

type service struct {
	Name     string
	Endpoint string
}

func getService(serviceName string) (service, error) {

	for _, service := range getServices() {
		if service.Name == serviceName {
			return service, nil
		}
	}
	return service{}, fmt.Errorf("%s service not found", serviceName)
}

func getServiceNameAndPath(path string) (string, string) {
	log.Println("path: ", path)
	for _, service := range getServices() {
		prefix := "/" + service.Name
		if strings.HasPrefix(path, prefix) {
			return service.Name, strings.TrimPrefix(path, prefix)
		}
	}
	return "", path
}

func getServices() [1]service {
	services := [...]service{{Name: "user", Endpoint: config.GetConfiguration().USER_SERVICE_ENDPOINT}}
	return services
}
