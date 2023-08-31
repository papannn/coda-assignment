package repository

import "github.com/papannn/coda-assignment/discovery-service/domain"

type IServiceRepository interface {
	GetNamespaceList() []string
	GetServiceListByNamespace(namespace string) (*domain.ServiceList, error)
	GetAvailableServiceCountByNamespace(namespace string) (int64, error)
	AddServiceByNamespace(namespace string, service domain.Service) error
	RemoveServiceByNamespace(namespace string, service domain.Service) error
}
