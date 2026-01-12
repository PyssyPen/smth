package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/chats")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		os.Exit(1)
	}

	fmt.Println("Получено:")
	fmt.Println(string(body))
}
