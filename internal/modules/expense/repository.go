package expense

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(exp *Expense) error
}

type repository struct { db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repository{db: db} }

func (r *repository) Create(exp *Expense) error { return r.db.Create(exp).Error }
