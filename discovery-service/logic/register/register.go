package register

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic"
)

type IRegister interface {
	Register(request api.RegisterRequest) error
}

type Impl struct {
	ServiceMap logic.ServiceMap
}

func (impl *Impl) Register(request api.RegisterRequest) error {
	// TODO write validation so no duplicate IP + Port will be in the list

	serviceList, ok := impl.ServiceMap[request.Namespace]
	if ok {
		serviceList.Services = append(serviceList.Services, logic.Service{
			IP:       request.IP,
			Port:     request.Port,
			IsActive: true,
		})
	} else {
		impl.ServiceMap[request.Namespace] = &logic.ServiceList{
			Services: []logic.Service{
				{
					IP:       request.IP,
					Port:     request.Port,
					IsActive: true,
				},
			},
		}
	}

	return nil
}
