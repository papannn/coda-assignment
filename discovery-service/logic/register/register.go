package register

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic"
)

func Register(request api.RegisterRequest) error {
	// TODO write validation so no duplicate IP + Port will be in the list

	serviceList, ok := logic.ServiceMap[request.Namespace]
	if ok {
		serviceList.Services = append(serviceList.Services, logic.Service{
			IP:   request.IP,
			Port: request.Port,
		})
	} else {
		logic.ServiceMap[request.Namespace] = &logic.ServiceList{
			Services: []logic.Service{
				{
					IP:   request.IP,
					Port: request.Port,
				},
			},
		}
	}

	return nil
}
