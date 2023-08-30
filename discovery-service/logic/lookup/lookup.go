package lookup

import (
	"errors"
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic"
	"github.com/papannn/coda-assignment/discovery-service/logic/load_balancer"
)

type ILookup interface {
	Lookup(req api.LookupRequest) (*api.LookupResponse, error)
}

type Impl struct {
	ServiceMap             logic.ServiceMap
	LoadBalancingAlgorithm load_balancer.ILoadBalancer
}

func (impl *Impl) Lookup(req api.LookupRequest) (*api.LookupResponse, error) {
	serviceList, ok := impl.ServiceMap[req.Namespace]
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

	service := impl.LoadBalancingAlgorithm.LoadBalance(serviceList)
	if !service.IsActive {
		return nil, errors.New("no service available at the moment")
	}

	return &api.LookupResponse{
		IP:   service.IP,
		Port: service.Port,
	}, nil
}
