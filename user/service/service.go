package service

import (
	"net/http"
	"strconv"
	"user/entity"
	"user/helper"
	"user/repository"

	"github.com/labstack/echo/v4"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo}
}

func (us *Service) Register(c echo.Context) error {
	var payload entity.RegisterUserPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	_, err := us.repo.Register(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
	})

}

func (us *Service) Login(c echo.Context) error {
	var payload entity.LoginUserPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	user, err := us.repo.Login(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	userId := strconv.Itoa(int(user.ID))

	token, err := helper.GenerateJWTToken(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"token":   token,
	})
}
