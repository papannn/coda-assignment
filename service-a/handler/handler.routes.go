package handler

import "net/http"

func (service *ServiceA) RegisterRoutes() {
	http.HandleFunc("/api/example", service.ExampleEndpoint)
	http.HandleFunc("/health_check", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
}
