package unregister

import "github.com/papannn/coda-assignment/discovery-service/api"

type IUnregister interface {
	Unregister(req api.UnregisterRequest) error
}
