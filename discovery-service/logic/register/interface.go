package register

import "github.com/papannn/coda-assignment/discovery-service/api"

type IRegister interface {
	Register(request api.RegisterRequest) error
}
