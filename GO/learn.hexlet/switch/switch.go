package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ModifySpaces(s, mode string) string {
	r := ""
	switch mode {
	case "dash":
		r = "-"
	case "underscore":
		r = "_"
	case "", "unknown":
		r = "*"
	}
	return strings.ReplaceAll(s, " ", r)
}

func main() {
	mode := ""
	fmt.Print("Введите строку: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print("Введите модификатор: ")
	fmt.Scanln(&mode)
	//mode, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	//fmt.Println(mode)
	r := ModifySpaces(s, mode)
	fmt.Print(r)
}
