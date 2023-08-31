package handler

import "net/http"

func (service *ServiceA) RegisterRoutes() {
	http.HandleFunc("/api/example", service.ExampleEndpoint)
}
