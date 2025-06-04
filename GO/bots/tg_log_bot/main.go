package main

import (
	"fmt"
	"log"
	"tg_log_bot/bot"
	"tg_log_bot/key"
)

func main() {
	token, err := key.ReadKey("/home/pyssy/VSC/GO/bots/tg_log_bot/key/key.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := bot.StartBot(token); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}
