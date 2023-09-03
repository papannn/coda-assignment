package health_check

import (
	"fmt"
	"github.com/papannn/coda-assignment/discovery-service/config"
	"github.com/papannn/coda-assignment/discovery-service/domain"
	"github.com/papannn/coda-assignment/discovery-service/repository"
	"log"
	"net/http"
	"time"
)

type Impl struct {
	Config     config.Config
	Repository repository.IServiceRepository
}

func (impl *Impl) HealthCheck() error {
	namespaceList := impl.Repository.GetNamespaceList()

	for _, namespace := range namespaceList {
		serviceList, err := impl.Repository.GetServiceListByNamespace(namespace)
		if err != nil {
			return err
		}

		for _, service := range serviceList.Services {
			impl.hitHealthCheck(namespace, service)
		}
	}

	return nil
}

func (impl *Impl) hitHealthCheck(namespace string, service *domain.Service) {
	client := http.Client{
		Timeout: time.Millisecond * time.Duration(impl.Config.TimeoutThreshold),
	}
	URL := fmt.Sprintf("http://%s:%s/health_check", service.IP, service.Port)
	req, _ := http.NewRequest(http.MethodGet, URL, nil)

	_, err := client.Do(req)
	if err != nil {
		log.Println(err)
		_ = impl.Repository.RemoveServiceByNamespace(namespace, *service)
	}
}
