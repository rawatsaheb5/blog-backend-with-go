package expensesplit

import "gorm.io/gorm"

type Repository interface {
	BulkCreate(splits []ExpenseSplit) error
}

type repository struct { db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repository{db: db} }

func (r *repository) BulkCreate(splits []ExpenseSplit) error {
	if len(splits) == 0 { return nil }
	return r.db.Create(&splits).Error
}
