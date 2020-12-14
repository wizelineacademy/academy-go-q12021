package services

//This file manages the dependency injection system

import (
	"fmt"
	"log"
)

//ServiceRegistry is the structure of the ServicesRegistry
type ServiceRegistry struct {
	services     map[string]Service
	serviceTypes []string
}

//ServicesRegistry is the global variable for store and fetch services
var ServicesRegistry *ServiceRegistry

//NewServiceRegistry returns the instance of the ServiceRegistry structure.
func NewServiceRegistry() *ServiceRegistry {
	ServicesRegistry = &ServiceRegistry{
		services: make(map[string]Service),
	}
	return ServicesRegistry
}

const (
	STATUSOK     = iota
	STATUSSTOPED = iota
)

//Service is the interface that an object needs to implement to be considered as a service
type Service interface {
	Start() error
	Stop() error
	Status() int
}

//RegisterService registers a new service in the  ServiceRegistry
func (s *ServiceRegistry) RegisterService(serviceName string, service Service) error {
	if _, exists := s.services[serviceName]; exists {
		return fmt.Errorf("Service already exists: %v", serviceName)
	}
	s.services[serviceName] = service
	s.serviceTypes = append(s.serviceTypes, serviceName)
	return nil
}

//StartAll starts all the services contained in the ServiceRegistry
/*func (s *ServiceRegistry) StartAll() {
	log.Printf("Starting %d services: %v\n", len(s.serviceTypes), s.serviceTypes)
	for _, kind := range s.serviceTypes {
		log.Printf("Starting service type %v\n", kind)
		go s.services[kind].Start()
	}
}*/

// StopAll ends every service in reverse order of registration, logging a
// panic if any of them fail to stop.
/*func (s *ServiceRegistry) StopAll() {
	for i := len(s.serviceTypes) - 1; i >= 0; i-- {
		kind := s.serviceTypes[i]
		service := s.services[kind]
		if err := service.Stop(); err != nil {
			log.Panicf("Could not stop the following service: %v, %v", kind, err)
		}
	}
}*/

// FetchService takes in a struct pointer and sets the value of that pointer
// to a service currently stored in the service registry. This ensures the input argument is
// set to the right pointer that refers to the originally registered service.
func (s *ServiceRegistry) FetchService(serviceType string) Service {
	if service, ok := s.services[serviceType]; ok {
		return service
	}
	log.Println("Requested service type no exists: " + serviceType)
	return nil
}
