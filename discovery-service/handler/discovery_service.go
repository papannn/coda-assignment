package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/config"
	"github.com/papannn/coda-assignment/discovery-service/logic/lookup"
	"github.com/papannn/coda-assignment/discovery-service/logic/register"
	"github.com/papannn/coda-assignment/discovery-service/logic/status"
	"github.com/papannn/coda-assignment/discovery-service/logic/unregister"
)

type DiscoveryService struct {
	Config          config.Config
	LookupLogic     lookup.ILookup
	RegisterLogic   register.IRegister
	UnregisterLogic unregister.IUnregister
	StatusLogic     status.IStatus
}
