package parser

import (
	"encoding/json"
	"io"
)

func ParseJSONBody(body io.ReadCloser, target any) error {
	requestByte, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(requestByte, target)
	if err != nil {
		return err
	}

	return nil
}
