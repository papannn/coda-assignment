package response_writer

import (
	"encoding/json"
	"net/http"
)

func Write(writer http.ResponseWriter, obj any, httpStatus int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatus)
	byteData, _ := json.Marshal(obj)
	writer.Write(byteData)
}
