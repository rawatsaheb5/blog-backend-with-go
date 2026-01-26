package user

import (
	"gorm.io/gorm"
)

type Repository interface{
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository{
	return &repository{db}
}

func (r *repository) CreateUser(user *User) error{
	return r.db.Create(user).Error
}

func (r *repository) GetUserByID(id string) (*User, error){
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *repository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *repository) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}