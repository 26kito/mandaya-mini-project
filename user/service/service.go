package service

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	if err := us.validateRegisterPayload(payload); err != nil {
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

	if err := us.validateLoginPayload(payload); err != nil {
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

func (us *Service) validateRegisterPayload(payload entity.RegisterUserPayload) error {
	if payload.FullName == "" {
		return fmt.Errorf("Full name is required")
	}

	if payload.NIK == "" {
		return fmt.Errorf("NIK is required")
	}

	if payload.Email == "" {
		return fmt.Errorf("Email is required")
	}

	if len(payload.Email) < 5 || len(payload.Email) > 30 {
		return fmt.Errorf("Email is invalid")
	}

	if !strings.Contains(payload.Email, "@") || !strings.Contains(payload.Email, ".com") {
		return fmt.Errorf("Email is invalid")
	}

	if payload.Password == "" {
		return fmt.Errorf("Password is required")
	}

	return nil
}

func (us *Service) validateLoginPayload(payload entity.LoginUserPayload) error {
	if payload.Email == "" {
		return fmt.Errorf("Email is required")
	}

	if payload.Password == "" {
		return fmt.Errorf("Password is required")
	}

	return nil
}
