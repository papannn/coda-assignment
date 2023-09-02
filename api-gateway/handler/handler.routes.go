package handler

import "net/http"

func (service *APIGateway) RegisterRoutes() {
	http.HandleFunc("/api/example", service.ExampleEndpoint)
}
