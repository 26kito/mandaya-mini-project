package main

import (
	"hotel/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.Routes()
}
