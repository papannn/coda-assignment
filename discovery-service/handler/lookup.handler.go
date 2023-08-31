package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"net/http"
)

func (service *DiscoveryService) LookupEndpoint(writer http.ResponseWriter, request *http.Request) {
	req := api.LookupRequest{}
	err := parser.ParseJSONBody(request.Body, &req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	resp, err := service.LookupLogic.Lookup(req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response_writer.Write(writer, resp, http.StatusOK)
}
