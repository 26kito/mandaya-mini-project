package repository

import (
	"fmt"
	"hotel/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetHotelList() (*[]entity.Hotel, error)
	GetHotelByID(id int) (*entity.HotelRoom, error)
	GetRoomDetail(payload entity.CheckRoomAvailabilityPayload) (*entity.Room, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetHotelList() (*[]entity.Hotel, error) {
	var hotels []entity.Hotel

	result := r.db.Find(&hotels)
	if result.Error != nil {
		return nil, result.Error
	}

	return &hotels, nil
}

func (r *repository) GetHotelByID(id int) (*entity.HotelRoom, error) {
	var hotel entity.HotelRoom

	result := r.db.Table("hotels").Where("id = ?", id).Preload("Rooms").First(&hotel)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("hotel with ID %d not found", id)
		}
		return nil, result.Error
	}

	return &hotel, nil
}

func (r *repository) GetRoomDetail(payload entity.CheckRoomAvailabilityPayload) (*entity.Room, error) {
	var room entity.Room

	result := r.db.Table("rooms").Where("hotel_id = ? AND id = ?", payload.HotelID, payload.RoomID).First(&room)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			fmt.Println("room not found")
			return nil, fmt.Errorf("room with ID %d not found", payload.RoomID)
		}
		return nil, result.Error
	}

	return &room, nil
}
