package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (i Implementation) Register(w http.ResponseWriter, r *http.Request) {
	// Извлечение имени сервиса из заголовка
	serviceName := r.Header.Get("X-Service-Name")
	if serviceName == "" {
		http.Error(w, "Service name header is missing", http.StatusBadRequest)
		return
	}

	// Получение пути к директории с сервисами из переменной окружения
	// TODO: move to config
	userServicesPath := os.Getenv("USER_SERVICES")
	if userServicesPath == "" {
		http.Error(w, "User services path not configured", http.StatusInternalServerError)
		return
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

	// Чтение и фильтрация тела запроса
	var requestData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	filteredData := make(map[string]interface{})
	for field := range fields {
		if value, exists := requestData[field]; exists {
			filteredData[field] = value
		}
	}

	// Кодирование отфильтрованных данных
	filteredBody, err := json.Marshal(filteredData)
	if err != nil {
		http.Error(w, "Failed to encode filtered data", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(filteredBody))

	// Отправка запроса на эндпоинт
	req, err := http.NewRequest("POST", url, bytes.NewReader(filteredBody))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create request: %s", err), http.StatusInternalServerError)
		return
	}

	// asdasd, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to read request body: %s", err), http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println(string(asdasd))

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send request: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Возврат ответа клиенту
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
