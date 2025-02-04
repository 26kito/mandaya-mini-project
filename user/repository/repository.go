package repository

import (
	"user/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Register(payload entity.RegisterUserPayload) (*entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Register(payload entity.RegisterUserPayload) (*entity.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	user := entity.User{
		FullName: payload.FullName,
		NIK:      payload.NIK,
		Email:    payload.Email,
		Password: string(hashedPassword),
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
