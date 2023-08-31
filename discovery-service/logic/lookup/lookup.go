package lookup

import (
	"errors"
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic/load_balancer"
	"github.com/papannn/coda-assignment/discovery-service/repository"
)

type ILookup interface {
	Lookup(req api.LookupRequest) (*api.LookupResponse, error)
}

type Impl struct {
	LoadBalancingAlgorithm load_balancer.ILoadBalancer
	Repository             repository.IServiceRepository
}

func (impl *Impl) Lookup(req api.LookupRequest) (*api.LookupResponse, error) {
	serviceList, err := impl.Repository.GetServiceListByNamespace(req.Namespace)
	if err != nil {
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
			IP:                    service.IP,
			Port:                  service.Port,
			ServiceAvailableCount: 1,
		}, nil
	}

	service := impl.LoadBalancingAlgorithm.LoadBalance(serviceList)
	if !service.IsActive {
		return nil, errors.New("no service available at the moment")
	}

	serviceCount, err := impl.Repository.GetAvailableServiceCountByNamespace(req.Namespace)
	if err != nil {
		return nil, err
	}

	return &api.LookupResponse{
		IP:                    service.IP,
		Port:                  service.Port,
		ServiceAvailableCount: serviceCount,
	}, nil
}
