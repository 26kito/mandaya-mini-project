package main

import (
	"reservation/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.Routes()
}
