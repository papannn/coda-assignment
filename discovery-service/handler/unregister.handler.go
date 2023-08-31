package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"net/http"
)

func (service *DiscoveryService) UnregisterEndpoint(writer http.ResponseWriter, request *http.Request) {
	req := api.UnregisterRequest{}
	err := parser.ParseJSONBody(request.Body, &req)
	if err != nil {
		response_writer.Write(writer, api.UnregisterResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = service.UnregisterLogic.Unregister(req)
	if err != nil {
		response_writer.Write(writer, api.UnregisterResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	response_writer.Write(writer, api.UnregisterResponse{
		Message: "success unregister",
	}, http.StatusOK)
}
