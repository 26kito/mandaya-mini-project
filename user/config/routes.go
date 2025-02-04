package config

import (
	"log"
	"user/repository"
	"user/service"

	"github.com/labstack/echo/v4"
)

func Routes() {
	// Initialize database
	db, err := InitDatabase()
	if err != nil {
		log.Println(err)
	}

	// Initialize repository
	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	// Initialize HTTP server
	e := echo.New()

	e.POST("/register", service.Register)

	e.Logger.Fatal(e.Start(":8080"))
}
