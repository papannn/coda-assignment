package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/papannn/coda-assignment/discovery-service/handler"
	health_check2 "github.com/papannn/coda-assignment/discovery-service/logic/health_check"
	"github.com/papannn/coda-assignment/discovery-service/logic/load_balancer"
	"github.com/papannn/coda-assignment/discovery-service/logic/lookup"
	"github.com/papannn/coda-assignment/discovery-service/logic/register"
	"github.com/papannn/coda-assignment/discovery-service/logic/status"
	"github.com/papannn/coda-assignment/discovery-service/logic/unregister"
	"github.com/papannn/coda-assignment/discovery-service/repository"
	"github.com/papannn/coda-assignment/discovery-service/repository/internal_var"
	"github.com/papannn/coda-assignment/discovery-service/scheduler/health_check"
	"github.com/papannn/coda-assignment/lib/config"
)

var (
	repositoryImpl         repository.IServiceRepository
	loadBalancingAlgorithm load_balancer.ILoadBalancer
)

func Serve() {
	app := handler.DiscoveryService{}

	config.ReadConfig(&app.Config)
	initiateGlobalVar(&app)
	injectLogic(&app)
	app.RegisterRoutes()

	go registerJobs(&app)

	log.Println(fmt.Sprintf("Running on address: %s:%s", app.Config.IP, app.Config.Port))
	http.ListenAndServe(fmt.Sprintf("%s:%s", app.Config.IP, app.Config.Port), nil)
}

func registerJobs(app *handler.DiscoveryService) {
	for range time.Tick(time.Second * time.Duration(app.Config.HealthCheckTimeInterval)) {
		health_check.SchedulerHealthCheck(app)
	}
}

func injectLogic(app *handler.DiscoveryService) {
	injectRegisterLogic(app)
	injectUnregisterLogic(app)
	injectLookupLogic(app)
	injectStatusLogic(app)
	injectHealthCheckLogic(app)
}

func injectLookupLogic(app *handler.DiscoveryService) {
	app.LookupLogic = &lookup.Impl{
		Repository:             repositoryImpl,
		LoadBalancingAlgorithm: loadBalancingAlgorithm,
	}
}

func injectRegisterLogic(app *handler.DiscoveryService) {
	app.RegisterLogic = &register.Impl{
		Repository: repositoryImpl,
	}
}

func injectUnregisterLogic(app *handler.DiscoveryService) {
	app.UnregisterLogic = &unregister.Impl{
		Repository: repositoryImpl,
	}
}

func injectStatusLogic(app *handler.DiscoveryService) {
	app.StatusLogic = &status.Impl{
		Repository: repositoryImpl,
	}
}

func injectHealthCheckLogic(app *handler.DiscoveryService) {
	app.HealthCheckLogic = health_check2.Impl{
		Repository: repositoryImpl,
		Config:     app.Config,
	}
}

func initiateGlobalVar(app *handler.DiscoveryService) {
	initiateLoadBalancingAlgorithm(app)
	initiateRepository(app)
}

func initiateRepository(app *handler.DiscoveryService) {
	repositoryImpl = internal_var.NewInternalVarImpl()
}

func initiateLoadBalancingAlgorithm(app *handler.DiscoveryService) {
	algorithm, ok := load_balancer.LoadBalancerMap[app.Config.LoadBalancingAlgorithm]
	if !ok {
		panic("please use the correct load balancing algorithm!")
	}

	loadBalancingAlgorithm = algorithm
}
