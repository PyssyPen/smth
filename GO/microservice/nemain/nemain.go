// package main

// // client.go (Микросервис 2)

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	resp, err := http.Get("http://localhost:8080/hello")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}

// 	fmt.Println("Response from Microservice 1:", string(body))
// }

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите сообщение для отправки:")

	// Считываем сообщение из стандартного ввода
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	resp, err := http.Post("http://localhost:8080/hello", "text/plain", strings.NewReader(message))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response from Microservice 1:", resp.Status)
}
