package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"project/internal/image"
	"project/internal/key"
	"project/internal/telegram"
	//"project/internal/server"
)

func main() {
	key, err := key.ReadKey("/home/pyssy/VSC/GO/bots/meme_bot/internal/key/key.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	imageService := image.NewImageService()
	telegramService := telegram.NewTelegramService(key, imageService)

	// создание сервера для работы http
	//httpServer := server.NewServer(imageService)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Print("Telegram servise starting...")
		telegramService.Start()
	}()

	// go func() {
	// 	httpServer.Start()
	// }()

	<-done

	log.Print("Telegram servise stopping...")
	telegramService.Stop()
	log.Print("Telegram servise stopped.")

}
