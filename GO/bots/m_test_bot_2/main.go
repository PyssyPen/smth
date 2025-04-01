package main

import (
	"fmt"
	"log"
	"m_test_bot_2/bot"
	"m_test_bot_2/pkg/key"
)

func main() {
	// Получаем токен бота из переменной окружения
	token, err := key.ReadKey("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/key/key.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = bot.StartBot(token)
	if err != nil {
		log.Fatal("Ошибка запуска бота:", err)
	}
}
