package repository

import "github.com/papannn/coda-assignment/discovery-service/domain"

type IServiceRepository interface {
	GetNamespaceList() []string
	GetServiceListByNamespace(namespace string) (*domain.ServiceList, error)
	AddServiceByNamespace(namespace string, service domain.Service) error
	RemoveServiceByNamespace(namespace string, service domain.Service) error
}
