package config

import (
	"log"
	"payment/middleware"
	"payment/repository"
	"payment/service"

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

	e.POST("/midtrans/payment", service.Payment, middleware.ValidateJWTMiddleware)

	e.Logger.Fatal(e.Start(":8083"))
}
