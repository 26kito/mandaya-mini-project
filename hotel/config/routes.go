package config

import (
	"hotel/repository"
	"hotel/service"
	"log"

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

	e.GET("/hotels", service.GetHotelList)
	e.GET("/hotels/:id", service.GetHotelByID)

	e.Logger.Fatal(e.Start(":8081"))
}
