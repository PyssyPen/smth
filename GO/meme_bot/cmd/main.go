package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"project/internal/image"
	//"project/internal/server"
	"project/internal/telegram"
)

func main() {
	imageService := image.NewImageService()
	telegramService := telegram.NewTelegramService("7446097555:AAGyOzHRUqwtrdQijoZX1gi-cWQXAS-ZX3Y", imageService)
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
