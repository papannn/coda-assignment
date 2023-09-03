package lookup

import "github.com/papannn/coda-assignment/discovery-service/api"

type ILookup interface {
	Lookup(req api.LookupRequest) (*api.LookupResponse, error)
}
