// // server.go
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// func main() {
// 	http.HandleFunc("/getmessage", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet {
// 			http.Error(w, "Ожидается метод GET", http.StatusMethodNotAllowed)
// 			return
// 		}
// 		fmt.Println("Клиент запросил сообщение.")
// 		fmt.Print("Напишите сообщение: ")

// 		// Считываем ввод пользователя из терминала
// 		reader := bufio.NewReader(os.Stdin)
// 		input, err := reader.ReadString('\n')
// 		if err != nil {
// 			http.Error(w, "Ошибка чтения сообщения", http.StatusInternalServerError)
// 			return
// 		}
// 		input = strings.TrimSpace(input)

// 		// Отсылаем сообщение клиенту
// 		fmt.Fprint(w, input)
// 		fmt.Println("Сообщение отправлено клиенту.")
// 	})

//		fmt.Println("Сервер запущен на порту 8081")
//		err := http.ListenAndServe(":8081", nil)
//		if err != nil {
//			fmt.Println("Ошибка запуска сервера:", err)
//		}
//	}
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	go startServer()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите сообщение (Б -> А): ")
		if !scanner.Scan() {
			break
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		err := sendMessage("http://localhost:8081/message", text)
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
		}
	}
}

func startServer() {
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		msg := new(bytes.Buffer)
		_, err := msg.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела", http.StatusInternalServerError)
			return
		}
		r.Body.Close()

		fmt.Printf("\n[Сообщение от А]: %s\n", msg.String())
		fmt.Print("Введите сообщение (Б -> А): ")
	})

	fmt.Println("Приложение Б слушает на :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func sendMessage(url, message string) error {
	resp, err := http.Post(url, "text/plain", bytes.NewBufferString(message))
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
