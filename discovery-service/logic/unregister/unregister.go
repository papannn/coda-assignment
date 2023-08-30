package unregister

import (
	"errors"
	"fmt"
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic"
	"slices"
)

func Unregister(req api.UnregisterRequest) error {
	serviceList, ok := logic.ServiceMap[req.Namespace]
	if !ok {
		return errors.New("namespace is not found")
	}

	resultIndex := -1
	size := len(serviceList.Services)
	for index, service := range serviceList.Services {
		if req.IP == service.IP && req.Port == service.Port {
			resultIndex = index
			break
		}
	}

	if resultIndex == -1 {
		return errors.New(fmt.Sprintf("IP + Port is not registered on namespace %s", req.Namespace))
	}

	serviceList.Services = slices.Delete(serviceList.Services, resultIndex, resultIndex+1)
	if resultIndex == size-1 {
		serviceList.Index = 0
	}
	return nil
}
