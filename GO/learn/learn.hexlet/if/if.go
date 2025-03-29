package main

import (
	"fmt"
	"os"
)

func DomainForLocale(domain, locale string) string {
	l := ""
	if locale == "" {
		l = "en"
	} else {
		l = locale
	}
	return fmt.Sprintf("%s.%s", l, domain)
	//fmt.Sprintf(l + "." + domain)
}

func main() {
	domain := ""
	locale := ""
	fmt.Print("Введите домен: ")
	fmt.Fscan(os.Stdin, &domain)
	fmt.Print("Введите язык (двумя буквами): ")
	//fmt.Fscan(os.Stdin, &locale)
	fmt.Scanln(&locale)
	r := DomainForLocale(domain, locale)
	fmt.Println(r)
}
