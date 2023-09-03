package status

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/repository"
)

type Impl struct {
	Repository repository.IServiceRepository
}

func (impl *Impl) Status() (*api.StatusResponse, error) {
	resp := make(map[string]api.ServiceList)

	namespaceList := impl.Repository.GetNamespaceList()

	for _, namespace := range namespaceList {
		var serviceListAPI []api.Service
		serviceList, err := impl.Repository.GetServiceListByNamespace(namespace)
		if err != nil {
			return nil, nil
		}

		for _, service := range serviceList.Services {
			serviceListAPI = append(serviceListAPI, api.Service{
				IP:       service.IP,
				Port:     service.Port,
				IsActive: service.IsActive,
			})
		}

		resp[namespace] = api.ServiceList{
			Services: serviceListAPI,
			Index:    serviceList.Index,
		}
	}

	return &api.StatusResponse{
		ServiceMap: resp,
	}, nil
}
