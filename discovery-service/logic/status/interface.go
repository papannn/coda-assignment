package status

import "github.com/papannn/coda-assignment/discovery-service/api"

type IStatus interface {
	Status() (*api.StatusResponse, error)
}
