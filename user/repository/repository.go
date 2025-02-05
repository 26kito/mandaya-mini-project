package repository

import (
	"fmt"
	"user/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Register(payload entity.RegisterUserPayload) (*entity.User, error)
	Login(payload entity.LoginUserPayload) (*entity.User, error)
	GetUserById(id int) (*entity.User, error)
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

	if err := r.validateRegister(payload); err != nil {
		return nil, err
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) Login(payload entity.LoginUserPayload) (*entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	return &user, nil
}

func (r *repository) GetUserById(id int) (*entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) validateRegister(payload entity.RegisterUserPayload) error {
	if err := r.db.Where("email = ?", payload.Email).First(&entity.User{}).Error; err == nil {
		return fmt.Errorf("Email already exists")
	}

	if err := r.db.Where("nik = ?", payload.NIK).First(&entity.User{}).Error; err == nil {
		return fmt.Errorf("NIK already exists")
	}

	return nil
}
