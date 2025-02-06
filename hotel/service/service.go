package service

import (
	"hotel/entity"
	"hotel/repository"
	"strconv"

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

func (s *Service) GetHotelByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	hotel, err := s.repo.GetHotelByID(id)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    hotel,
	})
}

func (s *Service) GetRoomDetail(c echo.Context) error {
	var payload entity.CheckRoomAvailabilityPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := s.repo.GetRoomDetail(payload)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    data,
	})
}
