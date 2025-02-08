package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"payment/entity"
	"payment/repository"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo}
}

func (s *Service) Payment(c echo.Context) error {
	// Get payload from request
	var payload entity.PaymentPayload

	// Get user ID from JWT
	getUserId := c.Get("user").(jwt.MapClaims)["user_id"].(string)

	userId, _ := strconv.Atoi(getUserId)

	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Bind payload to struct
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Fetch booking details by order ID
	getBookingDetail, err := s.repo.GetBookingByOrderID(payload.OrderID)
	booking := &getBookingDetail.Data
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if booking.GuestID != uint(userId) {
		return c.JSON(401, map[string]interface{}{
			"message": "Unauthorized access",
		})
	}

	// Validate booking status
	if err := s.validateBookingStatus(booking); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Retrieve user details
	getUserDetail, err := s.repo.GetUserProfile(booking.GuestID, tokenString)
	user := &getUserDetail.Data
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Prepare midtrans payload
	midtransPayload := s.prepareMidtransPayload(payload, user, booking.TotalPrice, "hotel booking")
	response, err := s.MidtransPaymentHandler(midtransPayload)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// Prepare payment method
	paymentMethod := response.PaymentType + " - " + response.VANumbers[0].Bank

	// Generate transaction ID
	transactionID := fmt.Sprintf("TRX-%d", time.Now().Unix())
	// Create and save payment entity
	payment := s.createPaymentEntity(transactionID, payload.OrderID, user.ID, booking.TotalPrice, "hotel booking", nil, "pending", paymentMethod)
	if err := s.repo.SavePayment(payment); err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	paymentResponse := entity.PaymentResponse{
		TransactionID:     transactionID,
		TransactionStatus: payment.PaymentStatus,
		Amount:            booking.TotalPrice,
		PaymentType:       payment.TransactionType,
		Bank:              response.VANumbers[0].Bank,
		VANumber:          response.VANumbers[0].VANumber,
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
		"data":    paymentResponse,
	})
}

func (s *Service) MidtransPaymentHandler(payload entity.MidtransPaymentPayload) (*entity.MidtransResponse, error) {
	var response entity.MidtransResponse

	client := resty.New()

	midtransServerKey := "SB-Mid-server-xeGDmSmVJV-RIlHxhJ4kX-b4"
	encodedKey := base64.StdEncoding.EncodeToString([]byte(midtransServerKey))

	url := "https://api.sandbox.midtrans.com/v2/charge"

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Basic %v", encodedKey)).
		SetBody(payload).
		Post(url)

	fmt.Println("Response:", resp)
	if err != nil {
		return nil, fmt.Errorf("500 | %v", err)
	}

	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		log.Println("Error unmarshalling response:", err)
		return nil, err
	}

	return &response, nil
}

func (s *Service) validateBookingStatus(booking *entity.DetailBooking) error {
	if booking.Status != "pending" {
		return fmt.Errorf("Booking has been %s", booking.Status)
	}

	return nil
}

func (s *Service) prepareMidtransPayload(payload entity.PaymentPayload, user *entity.UserDetail, amount float64, txName string) entity.MidtransPaymentPayload {
	return entity.MidtransPaymentPayload{
		PaymentType: "bank_transfer",
		TransactionDetail: struct {
			OrderID     string `json:"order_id"`
			GrossAmount string `json:"gross_amount"`
		}{
			OrderID:     payload.OrderID,
			GrossAmount: fmt.Sprintf("%.2f", amount),
		},
		CustomerDetail: struct {
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Phone     string `json:"phone"`
		}{
			Email:     user.Email,
			FirstName: user.FullName,
			LastName:  "",
			Phone:     "",
		},
		ItemDetails: []struct {
			ID       string `json:"id"`
			Price    string `json:"price"`
			Quantity int    `json:"quantity"`
			Name     string `json:"name"`
		}{
			{
				ID:       payload.OrderID,
				Price:    fmt.Sprintf("%.2f", amount),
				Quantity: 1,
				Name:     txName,
			},
		},
		BankTransfer: struct {
			Bank string `json:"bank"`
		}{
			Bank: payload.PaymentMethod,
		},
	}
}

func (s *Service) createPaymentEntity(transactionID string, orderID string, userID uint, amount float64, transactionType string, paymentDate *time.Time, paymentStatus string, paymentMethod string) entity.Payment {
	return entity.Payment{
		PaymentID:       transactionID,
		OrderID:         orderID,
		UserID:          userID,
		TotalAmount:     amount,
		TransactionType: transactionType,
		PaymentDate:     paymentDate,
		PaymentStatus:   paymentStatus,
		PaymentMethod:   paymentMethod,
	}
}
