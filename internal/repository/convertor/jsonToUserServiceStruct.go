package convertor

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	errMsgInvalidJson = "invalid JSON"
)

type Field struct {
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

type Endpoint struct {
	Path    string            `json:"path"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Fields  map[string]Field  `json:"fields"`
}

type Service struct {
	Name        string              `json:"name"`
	Url         string              `json:"url"`
	Description string              `json:"description"`
	Headers     map[string]string   `json:"headers"`
	Endpoints   map[string]Endpoint `json:"endpoints"`
}

func JSONToUserServiceStruct(pathToJSON string) (*Service, error) {
	var service *Service

	jsonBytes, err := os.ReadFile(pathToJSON)
	if err != nil {
		return nil, err
	}

	if !json.Valid(jsonBytes) {
		return nil, errors.New(errMsgInvalidJson)
	}

	err = json.Unmarshal(jsonBytes, service)
	if err != nil {
		return nil, err
	}

	return service, nil
}
