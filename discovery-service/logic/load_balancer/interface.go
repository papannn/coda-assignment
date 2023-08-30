package load_balancer

import "github.com/papannn/coda-assignment/discovery-service/logic"

type ILoadBalancer interface {
	LoadBalance(serviceList *logic.ServiceList) logic.Service
}