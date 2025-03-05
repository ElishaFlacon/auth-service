package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (s *service) ProcessRegistration(body []byte, serviceName string) ([]byte, error) {
	requestData, err := s.authRepository.ReadRequestData(body)
	if err != nil {
		return nil, err
	}

	fields, err := s.authRepository.GetFields(serviceName)
	if err != nil {
		return nil, err
	}

	filteredData := make(map[string]interface{})
	for field := range fields {
		if value, exists := requestData[field]; exists {
			filteredData[field] = value
		}
	}

	filteredBody, err := json.Marshal(filteredData)
	if err != nil {
		return nil, err
	}

	url, err := s.authRepository.GetServiceURL(serviceName)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(filteredBody))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
