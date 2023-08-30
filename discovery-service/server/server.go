package server

import (
	"fmt"
	"github.com/papannn/coda-assignment/discovery-service/config"
	"github.com/papannn/coda-assignment/discovery-service/handler"
	"log"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println(fmt.Sprintf("Running on address: %s:%s", config.ConfigObj.IP, config.ConfigObj.Port))

	http.ListenAndServe(fmt.Sprintf("%s:%s", config.ConfigObj.IP, config.ConfigObj.Port), mux)
}
