package group

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(group *Group) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(group *Group) error {
	return r.db.Create(group).Error
}
