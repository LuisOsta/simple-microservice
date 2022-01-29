package proxy

import "testing"

type GetServiceTest struct {
	serviceName   string
	expectSuccess bool
}

func TestGetService(t *testing.T) {
	getServiceTests := []GetServiceTest{
		{
			serviceName:   "profile",
			expectSuccess: true,
		},
		{
			serviceName:   "location",
			expectSuccess: false,
		},
	}

	for _, test := range getServiceTests {
		service, err := getService(test.serviceName)
		hasSucceeded := err == nil
		if hasSucceeded != test.expectSuccess {
			t.Errorf("Expected getService(%s) to %t, but it is %t", test.serviceName, test.expectSuccess, hasSucceeded)
		}

		if service.Name != test.serviceName && test.expectSuccess == true {
			t.Errorf("Expected service name to be %s, but it is %s", test.serviceName, service.Name)
		}
	}
}

type GetServiceNameAndPathTest struct {
	path         string
	expectedName string
	expectedPath string
}

func TestGetServiceNameAndPath(t *testing.T) {
	getServiceNameAndPathTests := []GetServiceNameAndPathTest{
		{
			path:         "/profile",
			expectedName: "profile",
			expectedPath: "",
		},
		{
			path:         "/profile/12345",
			expectedName: "profile",
			expectedPath: "/12345",
		},
	}

	for _, test := range getServiceNameAndPathTests {
		name, path := getServiceNameAndPath(test.path)
		if name != test.expectedName {
			t.Errorf("Expected service name to be %s, but it is %s", test.expectedName, name)
		}

		if path != test.expectedPath {
			t.Errorf("Expected service path to be %s, but it is %s", test.expectedPath, path)
		}
	}
}

func TestGetServices(t *testing.T) {
	expectedServices := map[string]bool{
		"profile": true,
	}

	for _, service := range getServices() {

		if !expectedServices[service.Name] {
			t.Errorf("Expected service %s not found", service.Name)
		}
	}
}
