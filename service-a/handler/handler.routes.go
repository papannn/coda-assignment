package handler

import (
	"net/http"
)

func (service *ServiceA) RegisterRoutes() {
	http.HandleFunc("/api/example", service.ExampleEndpoint)
	http.HandleFunc("/health_check", service.HealthCheckEndpoint)
}
