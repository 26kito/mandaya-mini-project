package entity

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type RegisterUserPayload struct {
	FullName string `json:"full_name"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
