package main

import (
	"fmt"
	"strings"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

var (
	invalidRequest = "invalid request"
)

func Validate(req UserCreateRequest) string {
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalidRequest
	}

	if req.Age <= 0 || req.Age > 150 {
		return invalidRequest
	}

	return ""
}

func main() {
	fn, a := "", 0
	fmt.Print("Введите имя: ")
	fmt.Scanln(&fn)
	fmt.Print("Введите возраст: ")
	fmt.Scanln(&a)
	u := UserCreateRequest{FirstName: fn, Age: a}
	fmt.Println("Ваше имя и возраст:", u.FirstName, u.Age)
	r := Validate(u)
	fmt.Println(r)
}
