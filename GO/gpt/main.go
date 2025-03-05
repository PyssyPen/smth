package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// ChatMessage структура для хранения сообщений чата
type ChatMessage struct {
	Username string
	Message  string
	IsBot    bool
}

// ChatData структура для передачи данных в шаблон
type ChatData struct {
	Messages []ChatMessage
}

var messages = []ChatMessage{}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Обработка отправки сообщения
		username := r.FormValue("username")
		message := r.FormValue("message")
		messages = append(messages, ChatMessage{Username: username, Message: message})

		// Генерация ответа с помощью OpenAI
		botResponse, err := generateBotResponse(message)
		if err != nil {
			log.Printf("Error generating bot response: %v", err)
			http.Error(w, "Failed to generate bot response", http.StatusInternalServerError)
			return
		}
		messages = append(messages, ChatMessage{Username: "Bot", Message: botResponse, IsBot: true})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		// Отображение страницы чата
		tmpl := template.Must(template.ParseFiles("chat.html"))
		data := ChatData{Messages: messages}
		tmpl.Execute(w, data)
	}
}

func generateBotResponse(message string) (string, error) {
	// Вставьте ваш OpenAI API ключ сюда
	openaiAPIKey := "sk-1234uvwxabcd5678uvwxabcd1234uvwxabcd5678"

	url := "https://api.openai.com/v1/chat/completions"
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to generate response: %s", resp.Status)
	}

	var responseBody struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	if len(responseBody.Choices) > 0 {
		return responseBody.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response generated")
}

func main() {
	// Настройка маршрутов
	http.HandleFunc("/", chatHandler)

	// Запуск сервера
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
