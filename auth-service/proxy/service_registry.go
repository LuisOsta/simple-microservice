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
		if strings.HasPrefix(path, service.Name) {
			return service.Name, "/" + strings.TrimPrefix(path, "/"+service.Name)
		}
	}
	return "", path
}

func getServices() [1]service {
	services := [...]service{{Name: "user", Endpoint: config.GetConfiguration().USER_SERVICE_ENDPOINT}}
	return services
}
