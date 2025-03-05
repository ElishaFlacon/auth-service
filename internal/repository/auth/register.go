package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func (r *repository) GetRegisterJSONData(serviceName string) (string, error) {
	// Получение пути к директории с сервисами из переменной окружения
	// TODO: move to config
	userServicesPath := os.Getenv("USER_SERVICES")
	if userServicesPath == "" {
		return "", fmt.Errorf("user services path not configured")
	}

	// Путь к файлу JSON
	filePath := filepath.Join(userServicesPath, fmt.Sprintf("%s.json", serviceName))

	// Чтение JSON-файла
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Service configuration not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Декодирование JSON
	var serviceConfig map[string]interface{}
	if err := json.NewDecoder(file).Decode(&serviceConfig); err != nil {
		http.Error(w, "Failed to decode service configuration", http.StatusInternalServerError)
		return
	}

	// Извлечение информации об эндпоинте
	endpoints, ok := serviceConfig["endpoints"].(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid service configuration", http.StatusInternalServerError)
		return
	}

	registerEndpoint, ok := endpoints["register"].(map[string]interface{})
	if !ok {
		http.Error(w, "Register endpoint not found", http.StatusInternalServerError)
		return
	}

	serviceUrl, ok := serviceConfig["url"].(string)
	if !ok {
		http.Error(w, "Service URL not found", http.StatusInternalServerError)
		return
	}

	endpointPath, ok := registerEndpoint["path"].(string)
	if !ok {
		http.Error(w, "Invalid register endpoint configuration", http.StatusInternalServerError)
		return
	}

	url := serviceUrl + endpointPath

	fields, ok := registerEndpoint["fields"].(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid fields configuration", http.StatusInternalServerError)
		return
	}
	return "", nil
}
