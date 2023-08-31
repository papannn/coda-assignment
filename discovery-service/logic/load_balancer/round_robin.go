package load_balancer

import (
	"github.com/papannn/coda-assignment/discovery-service/domain"
)

type RoundRobin struct {
}

func (algorithm *RoundRobin) LoadBalance(serviceList *domain.ServiceList) domain.Service {
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

	return serviceList.Services[serviceList.Index]
}
