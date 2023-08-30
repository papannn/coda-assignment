package handler

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/register", Register)
	mux.HandleFunc("/api/unregister", Unregister)
}
