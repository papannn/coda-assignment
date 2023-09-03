package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/papannn/coda-assignment/api-gateway/handler"
	"github.com/papannn/coda-assignment/api-gateway/logic/service_hit"
	"github.com/papannn/coda-assignment/lib/file"
)

func Serve() {
	app := handler.APIGateway{}
	file.ReadFile(&app.Config, "/config/config.json")
	app.RegisterRoutes()
	injectServiceHitImpl(&app)

	log.Println(fmt.Sprintf("Running on address: %s:%s", app.Config.IP, app.Config.Port))
	http.ListenAndServe(fmt.Sprintf("%s:%s", app.Config.IP, app.Config.Port), nil)
}

func injectServiceHitImpl(app *handler.APIGateway) {
	app.ServiceHitLogic = &service_hit.Impl{
		Config: app.Config,
	}
}
