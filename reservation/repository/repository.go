package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reservation/entity"
	"time"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type Repository interface {
	Reservation(userId int, payload entity.ReservationPayload) (*entity.Reservation, error)
	GetBookingByOrderID(orderID string) (*entity.Reservation, error)
	CheckIn(userId int, checkInPayload entity.CheckInPayload) (*entity.Reservation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Reservation(userId int, payload entity.ReservationPayload) (*entity.Reservation, error) {
	checkIn, err := time.Parse("2006-01-02", payload.CheckIn)
	if err != nil {
		return nil, err
	}

	checkOut, err := time.Parse("2006-01-02", payload.CheckOut)
	if err != nil {
		return nil, err
	}

	if err := r.validateReservation(checkIn, checkOut); err != nil {
		return nil, err
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"hotel_id": payload.HotelID,
			"room_id":  payload.RoomID,
		}).
		Post("http://localhost:8081/get-room-detail")

	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var room entity.RoomResponse
	if err := json.Unmarshal(resp.Body(), &room); err != nil {
		log.Fatal("Error parsing response:", err)
	}

	if room.Data.Status == "occupied" {
		return nil, errors.New("Room is not available")
	}

	// Calculate total price
	// Total days is the difference in time divided by 24 hours
	totalDays := int(checkOut.Sub(checkIn).Hours() / 24)
	totalPrice := (float64(payload.TotalGuest) * room.Data.Price) * float64(totalDays)

	bookingCode := fmt.Sprintf("BKNG-%v%v", userId, time.Now().Unix())

	reservation := entity.Reservation{
		BookingCode: bookingCode,
		GuestID:     uint(userId),
		HotelID:     room.Data.HotelID,
		RoomID:      room.Data.ID,
		CheckIn:     checkIn,
		CheckOut:    checkOut,
		TotalGuest:  payload.TotalGuest,
		TotalPrice:  totalPrice,
		Status:      "pending",
		IsCheckedIn: false,
	}

	if err := r.db.Create(&reservation).Error; err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *repository) GetBookingByOrderID(orderID string) (*entity.Reservation, error) {
	var booking entity.Reservation

	result := r.db.Where("booking_code = ?", orderID).First(&booking)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, fmt.Errorf("Booking code %s not found", orderID)
		}
		return nil, result.Error
	}

	return &booking, nil
}

func (r *repository) validateReservation(checkIn, checkOut time.Time) error {
	// Get today's date at midnight (ignoring time part)
	today := time.Now().Truncate(24 * time.Hour)

	if checkIn.Before(today) {
		return errors.New("Check in date must be greater than today")
	}

	if checkOut.Before(checkIn) {
		return errors.New("Check out date must be greater than check in date")
	}

	return nil
}

func (r *repository) CheckIn(userId int, checkInPayload entity.CheckInPayload) (*entity.Reservation, error) {
	bookingCode := checkInPayload.OrderID

	booking, err := r.GetBookingByOrderID(bookingCode)
	if err != nil {
		return nil, err
	}

	// Check if the booking belongs to the user
	if int(booking.GuestID) != userId {
		return nil, fmt.Errorf("Booking does not belong to the user")
	}

	// Check if the booking is already checked in
	if booking.IsCheckedIn {
		return nil, fmt.Errorf("Booking has already checked in")
	}

	// Update booking status
	booking.IsCheckedIn = true
	booking.Status = "checked-in"

	if err := r.db.Save(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}
