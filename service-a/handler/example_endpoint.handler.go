package handler

import (
	"fmt"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"github.com/papannn/coda-assignment/service-a/api"
	"net/http"
)

func (service *ServiceA) ExampleEndpoint(writer http.ResponseWriter, request *http.Request) {
	req := api.ExampleRequest{}
	err := parser.ParseJSONBody(request.Body, &req)
	if err != nil {
		response_writer.Write(writer, api.ExampleResponse{}, http.StatusBadRequest)
		return
	}

	writer.Header().Set("X-Forwarded-Host", fmt.Sprintf("%s : %s:%s", service.Config.Namespace, service.Config.IP, service.Config.Port))
	response_writer.Write(writer, req, http.StatusOK)
}