package entity

import "time"

type Payment struct {
	ID              uint       `gorm:"primaryKey;autoIncrement"`
	PaymentID       string     `gorm:"unique;not null" json:"payment_id"`
	OrderID         string     `gorm:"unique;not null" json:"order_id"`
	UserID          uint       `gorm:"not null" json:"user_id"`
	TotalAmount     float64    `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	TransactionType string     `gorm:"type:varchar(20);not null" json:"transaction_type"` // "topup" or "booking"
	PaymentDate     *time.Time `gorm:"type:date" json:"payment_date"`
	PaymentStatus   string     `gorm:"type:varchar(10);not null" json:"payment_status"`
	PaymentMethod   string     `gorm:"type:varchar(20);not null" json:"payment_method"`
	CreatedAt       time.Time  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"type:timestamp" json:"updated_at"`
}

type PaymentPayload struct {
	OrderID       string `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
}

type PaymentResponse struct {
	TransactionID     string  `json:"transaction_id"`
	TransactionStatus string  `json:"transaction_status"`
	Amount            float64 `json:"amount"`
	PaymentType       string  `json:"payment_type"`
	Bank              string  `json:"bank,omitempty"`
	VANumber          string  `json:"va_number,omitempty"`
}

type MidtransResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	Currency          string `json:"currency"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	ExpiryTime        string `json:"expiry_time"`
	VANumbers         []struct {
		Bank     string `json:"bank"`
		VANumber string `json:"va_number"`
	} `json:"va_numbers"`
}

type MidtransPaymentPayload struct {
	PaymentType       string `json:"payment_type"`
	TransactionDetail struct {
		OrderID     string `json:"order_id"`
		GrossAmount string `json:"gross_amount"`
	} `json:"transaction_details"`
	CustomerDetail struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
	} `json:"customer_details"`
	ItemDetails []struct {
		ID       string `json:"id"`
		Price    string `json:"price"`
		Quantity int    `json:"quantity"`
		Name     string `json:"name"`
	} `json:"item_details"`
	BankTransfer struct {
		Bank string `json:"bank"`
	} `json:"bank_transfer"`
}

type MidtransCallbackResponse struct {
	VA []struct {
		VANumber string `json:"va_number"`
		Bank     string `json:"bank"`
	} `json:"va_numbers"`
	TransactionTime   string        `json:"transaction_time"`
	TransactionStatus string        `json:"transaction_status"`
	TransactionID     string        `json:"transaction_id"`
	StatusMessage     string        `json:"status_message"`
	StatusCode        string        `json:"status_code"`
	SignatureKey      string        `json:"signature_key"`
	SettlementTime    string        `json:"settlement_time"`
	PaymentType       string        `json:"payment_type"`
	PaymentAmounts    []interface{} `json:"payment_amounts"`
	OrderID           string        `json:"order_id"`
	MerchantID        string        `json:"merchant_id"`
	GrossAmount       string        `json:"gross_amount"`
	FraudStatus       string        `json:"fraud_status"`
	Currency          string        `json:"currency"`
}

type GetDetailBookingResponse struct {
	Message string        `json:"message"`
	Data    DetailBooking `json:"data"`
}

type DetailBooking struct {
	ID          uint      `json:"id"`
	BookingCode string    `json:"booking_code"`
	GuestID     uint      `json:"guest_id"`
	HotelID     uint      `json:"hotel_id"`
	RoomID      uint      `json:"room_id"`
	CheckIn     time.Time `json:"check_in"`
	CheckOut    time.Time `json:"check_out"`
	TotalGuest  int       `json:"total_guest"`
	TotalPrice  float64   `json:"total_price"`
	Status      string    `json:"status"`
	IsCheckedIn bool      `json:"is_checked_in"`
}

type GetUserDetailResponse struct {
	Message string     `json:"message"`
	Data    UserDetail `json:"data"`
}

type UserDetail struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
}
