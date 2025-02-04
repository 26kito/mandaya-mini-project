package main

import (
	"log"
	"user/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Initialize database
	db, err := config.InitDatabase()
	if err != nil {
		log.Println(err)
	}
}
