package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Считываем тело запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Выводим полученное сообщение
	message := string(body)
	fmt.Println("Received message:", message)
	fmt.Println("Server shutdown")

	// Отправляем ответ клиенту
	fmt.Fprintln(w, "Message received by Microservice 1!")

	// Завершаем работу сервера после обработки запроса
	os.Exit(0)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting Microservice 1 on :8080")

	// Создаем канал для получения сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		fmt.Println("Shutting down Microservice 1...")
		os.Exit(0)
	}()

	http.ListenAndServe(":8080", nil)
}

//package main

// main.go (Микросервис 1)

// import (
// 	"fmt"
// 	"net/http"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello from Microservice 1!")
// }

// func main() {
// 	http.HandleFunc("/hello", helloHandler)
// 	fmt.Println("Starting Microservice 1 on :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// main.go (Микросервис 1)
