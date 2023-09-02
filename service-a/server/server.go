package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/papannn/coda-assignment/lib/config"
	"github.com/papannn/coda-assignment/service-a/handler"
)

func Serve() {
	app := handler.ServiceA{}
	config.ReadConfig(&app.Config)

	app.RegisterRoutes()

	for true {
		registerServiceOnStartup(app)
		log.Println(fmt.Sprintf("Running on address: %s:%d", app.Config.IP, app.Config.Port))
		_ = http.ListenAndServe(fmt.Sprintf("%s:%d", app.Config.IP, app.Config.Port), nil)
		oldPort := app.Config.Port
		app.Config.Port++
		log.Println(fmt.Sprintf("Port: %d already been used, trying to use port %d", oldPort, app.Config.Port))
	}

}

func registerServiceOnStartup(app handler.ServiceA) {
	payload := map[string]string{
		"namespace": app.Config.Namespace,
		"ip":        app.Config.IP,
		"port":      strconv.Itoa(app.Config.Port),
	}

	payloadByte, _ := json.Marshal(payload)

	URL := fmt.Sprintf("%s/api/register", app.Config.DiscoveryServiceBaseURL)

	http.Post(URL, "application/json", bytes.NewBuffer(payloadByte))
}

func unregisterService(app handler.ServiceA) {
	payload := map[string]string{
		"namespace": app.Config.Namespace,
		"ip":        app.Config.IP,
		"port":      strconv.Itoa(app.Config.Port),
	}

	payloadByte, _ := json.Marshal(payload)

	URL := fmt.Sprintf("%s/api/unregister", app.Config.DiscoveryServiceBaseURL)

	http.Post(URL, "application/json", bytes.NewBuffer(payloadByte))
}
