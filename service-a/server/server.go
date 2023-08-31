package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/papannn/coda-assignment/lib/config"
	"github.com/papannn/coda-assignment/service-a/handler"
	"net/http"
)

func Serve() {
	app := handler.ServiceA{}
	config.ReadConfig(&app.Config)

	registerServiceOnStartup(app)
}

func registerServiceOnStartup(app handler.ServiceA) {
	payload := map[string]string{
		"namespace": app.Config.Namespace,
		"ip":        app.Config.IP,
		"port":      app.Config.Port,
	}

	payloadByte, _ := json.Marshal(payload)

	URL := fmt.Sprintf("%s/api/register", app.Config.DiscoveryServiceBaseURL)

	http.Post(URL, "application/json", bytes.NewBuffer(payloadByte))
}
