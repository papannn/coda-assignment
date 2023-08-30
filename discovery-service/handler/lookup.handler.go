package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	logic "github.com/papannn/coda-assignment/discovery-service/logic/lookup"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"net/http"
)

func Lookup(writer http.ResponseWriter, request *http.Request) {
	req := api.LookupRequest{}
	err := parser.ParseRequest(request, &req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	resp, err := logic.Lookup(req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response_writer.Write(writer, resp, http.StatusOK)
}
