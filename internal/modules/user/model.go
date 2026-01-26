package user

import "time"

type User struct{
	ID string `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email string `gorm:"unique;not null"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}