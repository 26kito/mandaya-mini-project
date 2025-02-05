package repository

import (
	"hotel/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetHotelList() (*[]entity.Hotel, error)
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
