package server

import (
	"fmt"
	"github.com/papannn/coda-assignment/discovery-service/handler"
	"github.com/papannn/coda-assignment/discovery-service/logic"
	"github.com/papannn/coda-assignment/discovery-service/logic/load_balancer"
	"github.com/papannn/coda-assignment/discovery-service/logic/lookup"
	"github.com/papannn/coda-assignment/discovery-service/logic/register"
	"github.com/papannn/coda-assignment/discovery-service/logic/status"
	"github.com/papannn/coda-assignment/discovery-service/logic/unregister"
	"github.com/papannn/coda-assignment/lib/config"
	"log"
	"net/http"
)

var (
	serviceMap             logic.ServiceMap
	loadBalancingAlgorithm load_balancer.ILoadBalancer
)

func Serve() {
	app := handler.DiscoveryService{}

	config.ReadConfig(&app.Config)
	initiateGlobalVar(&app)
	injectLogic(&app)
	app.RegisterRoutes()

	log.Println(fmt.Sprintf("Running on address: %s:%s", app.Config.IP, app.Config.Port))
	http.ListenAndServe(fmt.Sprintf("%s:%s", app.Config.IP, app.Config.Port), nil)
}

func injectLogic(app *handler.DiscoveryService) {
	injectRegisterLogic(app)
	injectUnregisterLogic(app)
	injectLookupLogic(app)
	injectStatusLogic(app)
}

func injectLookupLogic(app *handler.DiscoveryService) {
	app.LookupLogic = &lookup.Impl{
		ServiceMap:             serviceMap,
		LoadBalancingAlgorithm: loadBalancingAlgorithm,
	}
}

func injectRegisterLogic(app *handler.DiscoveryService) {
	app.RegisterLogic = &register.Impl{
		ServiceMap: serviceMap,
	}
}

func injectUnregisterLogic(app *handler.DiscoveryService) {
	app.UnregisterLogic = &unregister.Impl{
		ServiceMap: serviceMap,
	}
}

func injectStatusLogic(app *handler.DiscoveryService) {
	app.StatusLogic = &status.Impl{
		ServiceMap: serviceMap,
	}
}

func initiateGlobalVar(app *handler.DiscoveryService) {
	serviceMap = make(map[string]*logic.ServiceList)
	initiateLoadBalancingAlgorithm(app)
}

func initiateLoadBalancingAlgorithm(app *handler.DiscoveryService) {
	algorithm, ok := load_balancer.LoadBalancerMap[app.Config.LoadBalancingAlgorithm]
	if !ok {
		panic("please use the correct load balancing algorithm!")
	}

	loadBalancingAlgorithm = algorithm
}
