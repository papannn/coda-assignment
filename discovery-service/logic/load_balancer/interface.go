package load_balancer

import (
	"github.com/papannn/coda-assignment/discovery-service/domain"
)

type ILoadBalancer interface {
	LoadBalance(serviceList *domain.ServiceList) domain.Service
}

var (
	LoadBalancerMap = map[string]ILoadBalancer{
		"round_robin": &RoundRobin{},
	}
)
