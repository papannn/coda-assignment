package handler

import (
	"net/http"
)

func (service *DiscoveryService) RegisterRoutes() {
	http.HandleFunc("/api/register", service.RegisterEndpoint)
	http.HandleFunc("/api/unregister", service.UnregisterEndpoint)
	http.HandleFunc("/api/lookup", service.LookupEndpoint)
	http.HandleFunc("/api/status", service.StatusEndpoint)
}
