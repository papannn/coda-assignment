package handler

import (
	"github.com/papannn/coda-assignment/api-gateway/config"
	"github.com/papannn/coda-assignment/api-gateway/logic/service_hit"
)

type APIGateway struct {
	Config          config.Config
	ServiceHitLogic service_hit.IServiceHit
}
