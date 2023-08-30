package status

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic"
)

type IStatus interface {
	Status() (*api.StatusResponse, error)
}

type Impl struct {
	ServiceMap logic.ServiceMap
}

func (impl *Impl) Status() (*api.StatusResponse, error) {
	resp := make(map[string]api.ServiceList)

	for key, list := range impl.ServiceMap {
		var serviceList []api.Service

		for _, service := range list.Services {
			serviceList = append(serviceList, api.Service{
				IP:       service.IP,
				Port:     service.Port,
				IsActive: service.IsActive,
			})
		}

		resp[key] = api.ServiceList{
			Services: serviceList,
			Index:    list.Index,
		}
	}

	return &api.StatusResponse{
		ServiceMap: resp,
	}, nil
}
