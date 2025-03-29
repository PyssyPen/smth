package main //hello

import (
	"fmt"
	"os"
	"strings"
)

func Printed() {
	fmt.Println("Hello, bitch")
}

func Greetings(n string) string {
	r := "Привет, " + strings.Title(strings.ToLower(strings.Trim(n, " "))) + "!"
	//r := strings.Trim(n, " ")
	//r = strings.ToLower(r)
	//r = strings.Title(r)
	//r = "Привет, " + r + "!"
	return r
}

func main() {
	Printed()
	var name string
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)
	r := Greetings(name)
	fmt.Println(r)
}
