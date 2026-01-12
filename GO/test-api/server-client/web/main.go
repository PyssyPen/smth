package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const apiAddr = "http://localhost:8000" // Адрес первого приложения (API-сервера)

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/books", proxyBooks)
	http.HandleFunc("/books/", proxyBooksByID)

	log.Println("Web-интерфейс стартовал на :8080 (фронтенд)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Отдаёт html-страницу клиенту
func serveIndex(w http.ResponseWriter, r *http.Request) {
	// можно встроить html прямо сюда или читать из файла
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Нет файла index.html", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Type", "text/html")
	io.Copy(w, file)
}

// Прокси для /books (GET, POST)
func proxyBooks(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, apiAddr+"/books")
}

// Прокси для /books/{id} (GET, PUT, DELETE)
func proxyBooksByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	proxyRequest(w, r, apiAddr+"/books/"+id)
}

// Универсальный проксирующий обработчик
func proxyRequest(w http.ResponseWriter, r *http.Request, url string) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, "Request creation error", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header.Clone()

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "API server unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Копируем заголовки и код ответа
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
