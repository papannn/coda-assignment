package unregister

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/domain"
	"github.com/papannn/coda-assignment/discovery-service/repository"
)

type IUnregister interface {
	Unregister(req api.UnregisterRequest) error
}

type Impl struct {
	Repository repository.IServiceRepository
}

func (impl *Impl) Unregister(req api.UnregisterRequest) error {
	err := impl.Repository.RemoveServiceByNamespace(req.Namespace, domain.Service{
		IP:   req.IP,
		Port: req.Port,
	})
	if err != nil {
		return err
	}

	return nil
}
