package handler

import "net/http"

func (service *ServiceA) HealthCheckEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
