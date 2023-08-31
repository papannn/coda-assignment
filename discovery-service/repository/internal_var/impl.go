package internal_var

import (
	"errors"
	"fmt"
	"github.com/papannn/coda-assignment/discovery-service/domain"
	"slices"
)

type ServiceMap = map[string]*domain.ServiceList

type Impl struct {
	ServiceMap ServiceMap
}

func NewInternalVarImpl() *Impl {
	return &Impl{
		ServiceMap: make(map[string]*domain.ServiceList),
	}
}

func (impl *Impl) GetNamespaceList() []string {
	var result []string

	for key, _ := range impl.ServiceMap {
		result = append(result, key)
	}

	return result
}

func (impl *Impl) GetServiceListByNamespace(namespace string) (*domain.ServiceList, error) {
	serviceList, ok := impl.ServiceMap[namespace]
	if !ok {
		return nil, errors.New("namespace is not found")
	}

	if len(serviceList.Services) == 0 {
		return nil, errors.New("namespace is not found")
	}

	return serviceList, nil
}

func (impl *Impl) AddServiceByNamespace(namespace string, service domain.Service) error {
	serviceList, ok := impl.ServiceMap[namespace]
	if ok {
		serviceList.Services = append(serviceList.Services, domain.Service{
			IP:       service.IP,
			Port:     service.Port,
			IsActive: true,
		})
	} else {
		impl.ServiceMap[namespace] = &domain.ServiceList{
			Services: []domain.Service{
				{
					IP:       service.IP,
					Port:     service.Port,
					IsActive: true,
				},
			},
		}
	}

	return nil
}

func (impl *Impl) RemoveServiceByNamespace(namespace string, service domain.Service) error {
	serviceList, ok := impl.ServiceMap[namespace]
	if !ok {
		return errors.New("namespace is not found")
	}

	resultIndex := -1
	size := len(serviceList.Services)
	for index, service := range serviceList.Services {
		if service.IP == service.IP && service.Port == service.Port {
			resultIndex = index
			break
		}
	}

	if resultIndex == -1 {
		return errors.New(fmt.Sprintf("IP + Port is not registered on namespace %s", namespace))
	}

	serviceList.Services = slices.Delete(serviceList.Services, resultIndex, resultIndex+1)
	if resultIndex == size-1 {
		serviceList.Index = 0
	}

	return nil
}

func (impl *Impl) GetAvailableServiceCountByNamespace(namespace string) (int64, error) {
	serviceList, err := impl.GetServiceListByNamespace(namespace)
	if err != nil {
		return 0, err
	}

	var count int64 = 0
	for _, service := range serviceList.Services {
		if service.IsActive {
			count++
		}
	}

	return count, nil
}
