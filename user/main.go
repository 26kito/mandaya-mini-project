package main

import (
	"user/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.Routes()
}
