package handler

import (
	"encoding/json"
	"github.com/papannn/coda-assignment/api"
	logic "github.com/papannn/coda-assignment/logic/register"
	"io"
	"net/http"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	// TODO add middleware for handler to minimize manually setting up the Content-Type and parsing the json req + resp
	writer.Header().Set("Content-Type", "application/json")
	requestByte, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	req := api.RegisterRequest{}
	err = json.Unmarshal(requestByte, &req)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		res := api.RegisterResponse{Message: "error unmarshal request body"}
		resByte, _ := json.Marshal(res)
		writer.Write(resByte)
		return
	}

	err = logic.Register(req)
	if err != nil {
		res := api.RegisterResponse{Message: err.Error()}
		resByte, _ := json.Marshal(res)
		writer.Write(resByte)
		return
	}
	res := api.RegisterResponse{Message: "success register service"}
	resByte, _ := json.Marshal(res)
	writer.Write(resByte)
}
