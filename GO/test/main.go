package main

import (
	"fmt"
)

// Пример функции, принимающей строку в качестве параметра
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// Вызов функции с передачей строки
	greet("Alice")
}
