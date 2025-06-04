package main

import (
	"log"

	"test/database"
	"test/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.InitDB()
	server.StartServer()
}
