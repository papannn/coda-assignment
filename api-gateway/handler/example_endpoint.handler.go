package handler

import (
	"github.com/papannn/coda-assignment/api-gateway/api"
	"github.com/papannn/coda-assignment/lib/parser"
	"github.com/papannn/coda-assignment/lib/response_writer"
	serviceAApi "github.com/papannn/coda-assignment/service-a/api"
	"net/http"
)

func (service *APIGateway) ExampleEndpoint(writer http.ResponseWriter, request *http.Request) {

	result, err := service.ServiceHitLogic.Post("service-a", request.RequestURI, request.Body)
	if err != nil {
		response_writer.Write(writer, api.ErrorMessageResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	resp := serviceAApi.ExampleResponse{}
	err = parser.ParseJSONBody(result.Body, &resp)
	if err != nil {
		response_writer.Write(writer, api.ErrorMessageResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("X-Forwarded-Host", result.Header.Get("X-Forwarded-Host"))
	response_writer.Write(writer, resp, http.StatusOK)
}
