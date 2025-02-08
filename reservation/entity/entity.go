package entity

import "time"

type Reservation struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	BookingCode string    `gorm:"type:varchar(30);not null" json:"booking_code"`
	GuestID     uint      `gorm:"not null" json:"guest_id"`
	HotelID     uint      `gorm:"not null" json:"hotel_id"`
	RoomID      uint      `gorm:"not null" json:"room_id"`
	CheckIn     time.Time `gorm:"type:date;not null" json:"check_in"`
	CheckOut    time.Time `gorm:"type:date;not null" json:"check_out"`
	TotalGuest  int       `gorm:"not null" json:"total_guest"`
	TotalPrice  float64   `gorm:"not null" json:"total_price"`
	Status      string    `gorm:"type:varchar(10);not null" json:"status"`
	IsCheckedIn bool      `gorm:"not null" json:"is_checked_in"`
}

type ReservationPayload struct {
	HotelID    int    `json:"hotel_id"`
	RoomID     int    `json:"room_id"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
	TotalGuest int    `json:"total_guest"`
}

type RoomResponse struct {
	Message string   `json:"message"`
	Data    RoomData `json:"data"`
}

type RoomData struct {
	ID         uint    `json:"id"`
	HotelID    uint    `json:"hotel_id"`
	RoomNumber string  `json:"room_number"`
	RoomType   string  `json:"room_type"`
	MaxGuest   int     `json:"max_guest"`
	Price      float64 `json:"price"`
	Status     string  `json:"status"`
}

type CheckInPayload struct {
	OrderID string `json:"order_id"`
}
