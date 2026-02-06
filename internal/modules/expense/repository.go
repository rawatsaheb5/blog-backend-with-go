package expense

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(exp *Expense) error
	ListByGroupID(groupID uint64) ([]Expense, error)
	GetByID(expenseID uint64) (*Expense, error)
}

type repository struct { db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repository{db: db} }

func (r *repository) Create(exp *Expense) error { return r.db.Create(exp).Error }

func (r *repository) ListByGroupID(groupID uint64) ([]Expense, error) {
	var exps []Expense
	if err := r.db.Where("group_id = ?", groupID).Order("expense_date DESC, id DESC").Find(&exps).Error; err != nil {
		return nil, err
	}
	return exps, nil
}

func (r *repository) GetByID(expenseID uint64) (*Expense, error) {
	var exp Expense
	if err := r.db.First(&exp, expenseID).Error; err != nil {
		return nil, err
	}
	return &exp, nil
}
