package handler

import (
	"net/http"

	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"github.com/papannn/coda-assignment/service-a/api"
)

func (service *ServiceA) ExampleEndpoint(writer http.ResponseWriter, request *http.Request) {
	var req interface{}
	err := parser.ParseJSONBody(request.Body, &req)
	if err != nil {
		response_writer.Write(writer, api.ExampleResponse{}, http.StatusBadRequest)
		return
	}

	response_writer.Write(writer, req, http.StatusOK)
}
