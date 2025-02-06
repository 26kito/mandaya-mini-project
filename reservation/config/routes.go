package config

import (
	"log"
	"reservation/middleware"
	"reservation/repository"
	"reservation/service"

	"github.com/labstack/echo/v4"
)

func Routes() {
	// Initialize database
	db, err := InitDatabase()
	if err != nil {
		log.Println(err)
	}

	e := echo.New()

	// Initialize repository
	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	e.POST("/reservation", service.Reservation, middleware.ValidateJWTMiddleware)

	e.Logger.Fatal(e.Start(":8082"))
}
