package entity

type Hotel struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"type:varchar(100);not null" json:"name"`
	Location      string `gorm:"type:varchar(255);not null" json:"location"`
	ContactNumber string `gorm:"type:varchar(15)" json:"contact_number"`
	Email         string `gorm:"type:varchar(100)" json:"email"`
}

type HotelRoom struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"type:varchar(100);not null" json:"name"`
	Location      string `gorm:"type:varchar(255);not null" json:"location"`
	ContactNumber string `gorm:"type:varchar(15)" json:"contact_number"`
	Email         string `gorm:"type:varchar(100)" json:"email"`
	Rooms         []Room `gorm:"foreignKey:HotelID" json:"rooms"`
}

type Room struct {
	ID         uint    `gorm:"primaryKey;autoIncrement"`
	HotelID    uint    `gorm:"not null" json:"hotel_id"`
	RoomNumber string  `gorm:"type:varchar(10);not null" json:"room_number"`
	RoomType   string  `gorm:"type:varchar(100);not null" json:"room_type"`
	MaxGuest   int     `gorm:"not null" json:"max_guest"`
	Price      float64 `gorm:"not null" json:"price"`
	Status     string  `gorm:"type:varchar(10);not null" json:"status"`
}
