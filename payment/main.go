package main

import (
	"payment/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.Routes()
}
