package service

import (
	"hotel/repository"

	"github.com/labstack/echo/v4"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetHotelList(c echo.Context) error {
	hotels, err := s.repo.GetHotelList()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    hotels,
	})
}
