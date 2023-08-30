package handler

import (
	"github.com/papannn/coda-assignment/discovery-service/api"
	logic "github.com/papannn/coda-assignment/discovery-service/logic/status"
	"github.com/papannn/coda-assignment/lib/response_writer"
	"net/http"
)

func Status(writer http.ResponseWriter, request *http.Request) {
	resp, err := logic.Status()
	if err != nil {
		response_writer.Write(writer, api.RegisterResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response_writer.Write(writer, resp, http.StatusOK)
}
