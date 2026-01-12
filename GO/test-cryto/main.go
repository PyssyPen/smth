package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "mytestpassword"
	// Генерируем хеш (bcrypt автоматически использует соль)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hash: %s\n", hash)
	fmt.Println("Исходный пароль:", password)

	// Теперь попробуем проверить пароль (это единственный способ!)
	check := "mytestpassword" // подставьте любой другой пароль для проверки

	err = bcrypt.CompareHashAndPassword(hash, []byte(check))
	if err != nil {
		fmt.Println("Проверка не пройдена: пароль неверный.")
	} else {
		fmt.Println("Проверка пройдена: текст пароля совпал с хешем!")
	}

	// ПОПЫТКИ восстановить пароль из ХЭША не будет, потому что это невозможно!
}
