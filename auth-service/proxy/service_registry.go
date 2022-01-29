package proxy

import (
	"fmt"
	"strings"

	"github.com/auth-service/config"
)

type service struct {
	Name     string
	Endpoint string
}

// If the service name is not found in the hashmap, then return an error.
// Otherwise return the service object.
func getService(serviceName string) (service, error) {

	services := getServices()

	if services[serviceName] == (service{}) {
		return service{}, fmt.Errorf("service %s not found", serviceName)
	} else {
		return services[serviceName], nil
	}
}

// Uses the full request path to determine the service that it is trying to target and the endpoint for that service that it is trying to reach.
func getServiceNameAndPath(path string) (string, string) {
	for _, service := range getServices() {
		prefix := "/" + service.Name
		if strings.HasPrefix(path, prefix) {
			return service.Name, strings.TrimPrefix(path, prefix)
		}
	}
	return "", path
}

func getServices() map[string]service {
	return map[string]service{
		"profile": {
			Name:     "profile",
			Endpoint: config.GetConfiguration().USER_SERVICE_ENDPOINT,
		},
	}
}
