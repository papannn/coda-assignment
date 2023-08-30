package lookup

import (
	"errors"
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic"
)

func Lookup(req api.LookupRequest) (*api.LookupResponse, error) {
	serviceList, ok := logic.ServiceMap[req.Namespace]
	if !ok {
		return nil, errors.New("namespace is not found")
	}

	if len(serviceList.Services) == 0 {
		return nil, errors.New("no service available at the moment")
	}

	if len(serviceList.Services) == 1 {
		service := serviceList.Services[0]
		if !service.IsActive {
			return nil, errors.New("no service available at the moment")
		}

		return &api.LookupResponse{
			IP:   service.IP,
			Port: service.Port,
		}, nil
	}

	doneCycle := false
	index := serviceList.Index
	for serviceList.Index != index || !doneCycle {
		doneCycle = true
		serviceList.Index++
		if serviceList.Index == int64(len(serviceList.Services)) {
			serviceList.Index = 0
		}

		if serviceList.Services[serviceList.Index].IsActive {
			break
		}
	}

	service := serviceList.Services[serviceList.Index]
	if !service.IsActive {
		return nil, errors.New("no service available at the moment")
	}

	return &api.LookupResponse{
		IP:   service.IP,
		Port: service.Port,
	}, nil
}
