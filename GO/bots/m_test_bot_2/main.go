package main

import (
	"fmt"
	"log"
	"m_test_bot_2/bot"
	"m_test_bot_2/pkg/cart"
	"m_test_bot_2/pkg/db"
	"m_test_bot_2/pkg/key"
)

func main() {
	token, err := key.ReadKey("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/key/key.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Инициализация базы данных
	connStr, err := db.ReadDB("/home/pyssy/VSC/GO/bots/m_test_bot_2/pkg/db/db.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := cart.InitDB(connStr); err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	err = bot.StartBot(token)
	if err != nil {
		log.Fatal("Ошибка запуска бота:", err)
	}
}
