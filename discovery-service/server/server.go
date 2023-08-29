package server

import (
	"github.com/papannn/coda-assignment/handler"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	http.ListenAndServe(":4444", mux)
}
