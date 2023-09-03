package register

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/domain"
	"github.com/papannn/coda-assignment/discovery-service/repository"
)

type Impl struct {
	Repository repository.IServiceRepository
}

func (impl *Impl) Register(request api.RegisterRequest) error {
	err := impl.Repository.AddServiceByNamespace(request.Namespace, domain.Service{
		IP:   request.IP,
		Port: request.Port,
	})
	if err != nil {
		return err
	}

	return nil
}
