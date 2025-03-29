package main

import (
	"fmt"
)

const (
	OkMsg        = "OK"
	CancelledMsg = "CANCELLED"
)

const (
	OkCode = iota
	CancelledCode
	UnknownCode
)

func ErrorMessageToCode(msg string) int {
	switch msg {
	case OkMsg:
		return OkCode
	case CancelledMsg:
		return CancelledCode
	}

	return UnknownCode
}

func main() {
	ms := ""
	fmt.Println("Какая ошибка: ")
	fmt.Scanln(&ms)
	r := ErrorMessageToCode(ms)
	fmt.Println(r)
}

/*
func ErrorMessageToCode(msg string) int {
	r := msg
	n := 2
	switch r {
	case "OK":
		n = 0
	case "CANCELLED":
		n = 1
	case "UNKNOWN": //case "UNKNOWN", "err", "":
		n = 2
	}
	return n
}

func ErrorCodeToMessage(msg int) string {
	r := msg
	n := "UNKNOWN"
	switch r {
	case 0:
		n = "OK"
	case 1:
		n = "CANCELLED"
	case 2:
		n = "UNKNOWN"
	}
	return n
}
*/
