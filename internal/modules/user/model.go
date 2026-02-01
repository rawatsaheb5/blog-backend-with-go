package user

import "time"

type User struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `json:"name"`
	Email          string    `gorm:"unique;not null" json:"email"`
	HashedPassword string    `json:"hashed_password"`
	IsActive       bool      `json:"is_active"`
	ProfileUrl     string    `json:"profile_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
