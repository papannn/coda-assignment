package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	"github.com/papannn/coda-assignment/discovery-service/logic/register"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"net/http"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	req := api.RegisterRequest{}
	err := parser.ParseRequest(request, &req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	err = register.Register(req)
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response_writer.Write(writer, api.RegisterResponse{
		Message: "success register service",
	}, http.StatusOK)
}
