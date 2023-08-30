package parser

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequest(request *http.Request, target any) error {
	requestByte, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(requestByte, target)
	if err != nil {
		return err
	}

	return nil
}
