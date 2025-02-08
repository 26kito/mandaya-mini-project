package service

import (
	"fmt"
	"net/http"
	"reservation/entity"
	"reservation/repository"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Reservation(c echo.Context) error {
	// Get user_id from JWT
	getUserId := c.Get("user").(jwt.MapClaims)["user_id"].(string)

	// Convert string to int
	userId, err := strconv.Atoi(getUserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var payload entity.ReservationPayload
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	reservation, err := s.repo.Reservation(userId, payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    reservation,
	})
}

func (s *Service) GetBookingByOrderID(c echo.Context) error {
	orderID := c.Param("order_id")

	booking, err := s.repo.GetBookingByOrderID(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	fmt.Println(booking)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    booking,
	})
}

func (s *Service) CheckIn(c echo.Context) error {
	// Get user_id from JWT
	getUserId := c.Get("user").(jwt.MapClaims)["user_id"].(string)

	// Convert string to int
	userId, err := strconv.Atoi(getUserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var payload entity.CheckInPayload
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	reservation, err := s.repo.CheckIn(userId, payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    reservation,
	})
}
