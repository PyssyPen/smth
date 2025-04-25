package main

import (
	"fmt"
	"log"
	"m_test_bot_2/bot"
	"m_test_bot_2/pkg/cart"
	"m_test_bot_2/pkg/db"
	"m_test_bot_2/pkg/key"
	"m_test_bot_2/pkg/order"
	"m_test_bot_2/pkg/profile"
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
		log.Fatal("Ошибка подключения к базе данных cart:", err)
	}
	if err := order.InitDB(connStr); err != nil {
		log.Fatal("Ошибка подключения к базе данных order:", err)
	}
	if err := profile.InitDB(connStr); err != nil {
		log.Fatal("Ошибка подключения к базе данных profile:", err)
	}

	err = bot.StartBot(token)
	if err != nil {
		log.Fatal("Ошибка запуска бота:", err)
	}
}
