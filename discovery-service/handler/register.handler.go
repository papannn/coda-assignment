package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"net/http"
)

func (service *DiscoveryService) RegisterEndpoint(writer http.ResponseWriter, request *http.Request) {
	req := api.RegisterRequest{}
	err := parser.ParseJSONBody(request.Body, &req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	err = service.RegisterLogic.Register(req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response_writer.Write(writer, api.RegisterResponse{
		Message: "success register service",
	}, http.StatusOK)
}
