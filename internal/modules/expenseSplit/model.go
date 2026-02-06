
package expensesplit

import "time"

type ExpenseSplit struct {
	// Composite primary key
	ExpenseID uint64 `gorm:"primaryKey" json:"expense_id"`
	UserID    uint64 `gorm:"primaryKey" json:"user_id"`

	// Amount user owes for this expense
	ShareAmount int64 `gorm:"not null" json:"share_amount"`

	CreatedAt time.Time `json:"created_at"`
}
