// // client.go
// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	url := "http://localhost:8081/getmessage"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Ошибка запроса:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Ошибка чтения ответа:", err)
// 		return
// 	}

//		message := string(body)
//		fmt.Println("Полученное сообщение от сервера:")
//		fmt.Println(message)
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

	// Читаем ввод из консоли, и отправляем сообщение на http://localhost:8082/message
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите сообщение (А -> Б): ")
		if !scanner.Scan() {
			break
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		err := sendMessage("http://localhost:8082/message", text)
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

		// Выводим полученное сообщение в консоль
		fmt.Printf("\n[Сообщение от Б]: %s\n", msg.String())
		fmt.Print("Введите сообщение (А -> Б): ") // Подсказка вновь
	})

	fmt.Println("Приложение А слушает на :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func sendMessage(url, message string) error {
	resp, err := http.Post(url, "text/plain", bytes.NewBufferString(message))
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
