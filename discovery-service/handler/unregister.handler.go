package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	logic "github.com/papannn/coda-assignment/discovery-service/logic/unregister"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"net/http"
)

func Unregister(writer http.ResponseWriter, request *http.Request) {
	req := api.UnregisterRequest{}
	err := parser.ParseRequest(request, &req)
	if err != nil {
		response_writer.Write(writer, api.UnregisterResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = logic.Unregister(req)
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
