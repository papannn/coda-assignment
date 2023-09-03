package internal_var

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/papannn/coda-assignment/lib/file"

	"github.com/papannn/coda-assignment/discovery-service/domain"
)

type ServiceMap = map[string]*domain.ServiceList

type Impl struct {
	ServiceMap ServiceMap
}

func NewInternalVarImpl() *Impl {
	var serviceMap ServiceMap
	err := file.ReadFile(&serviceMap, "/repository/internal_var/services.json")
	if err != nil {
		serviceMap = make(map[string]*domain.ServiceList)
	}
	return &Impl{
		ServiceMap: serviceMap,
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
		for _, currentService := range serviceList.Services {
			if service.IP == currentService.IP && service.Port == currentService.Port {
				currentService.IsActive = true
				impl.saveDataToJson()
				return nil
			}
		}

		serviceList.Services = append(serviceList.Services, &domain.Service{
			IP:       service.IP,
			Port:     service.Port,
			IsActive: true,
		})
	} else {
		impl.ServiceMap[namespace] = &domain.ServiceList{
			Services: []*domain.Service{
				{
					IP:       service.IP,
					Port:     service.Port,
					IsActive: true,
				},
			},
		}
	}

	impl.saveDataToJson()
	return nil
}

func (impl *Impl) RemoveServiceByNamespace(namespace string, currService domain.Service) error {
	serviceList, ok := impl.ServiceMap[namespace]
	if !ok {
		return errors.New("namespace is not found")
	}

	resultIndex := -1
	size := len(serviceList.Services)
	for index, service := range serviceList.Services {
		if currService.IP == service.IP && currService.Port == service.Port {
			resultIndex = index
			break
		}
	}

	if resultIndex == -1 {
		return errors.New(fmt.Sprintf("IP + Port is not registered on namespace %s", namespace))
	}

	serviceList.Services[resultIndex].IsActive = false
	if resultIndex == size-1 {
		serviceList.Index = 0
	}

	impl.saveDataToJson()
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

func (impl *Impl) saveDataToJson() {
	byteData, _ := json.Marshal(impl.ServiceMap)
	_ = os.WriteFile("repository/internal_var/services.json", byteData, 0644)
}
