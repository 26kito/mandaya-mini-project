package repository

import (
	"encoding/json"
	"fmt"
	"payment/entity"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type Repository interface {
	// Payment(userID int, payload entity.PaymentPayload) (*entity.PaymentResponse, error)
	GetBookingByOrderID(orderID string) (*entity.GetDetailBookingResponse, error)
	GetUserProfile(userID uint, tokenString string) (*entity.GetUserDetailResponse, error)
	SavePayment(payment entity.Payment) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (pr *repository) SavePayment(payment entity.Payment) error {
	result := pr.db.Create(&payment)

	if result.Error != nil {
		return fmt.Errorf("500 | %v", result.Error)
	}

	return nil
}

func (pr *repository) GetBookingByOrderID(orderID string) (*entity.GetDetailBookingResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("http://localhost:8082/reservation/" + orderID)
	if err != nil {
		return nil, fmt.Errorf("500 | %v", err)
	}

	var detailBooking entity.GetDetailBookingResponse

	if err := json.Unmarshal(resp.Body(), &detailBooking); err != nil {
		return nil, fmt.Errorf("500 | %v", err)
	}

	return &detailBooking, nil
}

func (pr *repository) GetUserProfile(userID uint, tokenString string) (*entity.GetUserDetailResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+tokenString).
		Get("http://localhost:8080/profile")
	if err != nil {
		return nil, fmt.Errorf("500 | %v", err)
	}

	var user entity.GetUserDetailResponse

	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, fmt.Errorf("500 | %v", err)
	}

	return &user, nil
}

// func (pr *repository) validateBookingStatus(booking *entity.GetDetailBookingResponse) error {
// 	if booking.Status != "pending" {
// 		return fmt.Errorf("400 | Booking has been %s", booking.Status)
// 	}

// 	return nil
// }
