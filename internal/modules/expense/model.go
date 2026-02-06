
package expense

import "time"

type Expense struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"expense_id"`

	// Group relation
	GroupID uint64 `gorm:"not null;index" json:"group_id"`

	// Expense details
	Title       string  `gorm:"type:text;not null" json:"title"`
	TotalAmount float64 `gorm:"not null" json:"total_amount"`

	// Payment info
	PaidBy uint64 `gorm:"not null;index" json:"paid_by"`

	// Expense metadata
	ExpenseDate time.Time `gorm:"not null" json:"expense_date"`
	SplitType   string    `gorm:"type:varchar(20);not null" json:"split_type"`
	Note        string    `gorm:"type:text" json:"note,omitempty"`

	// Audit
	CreatedBy uint64 `gorm:"not null;index" json:"created_by"`
	Status    string `gorm:"type:varchar(20);not null;default:'active'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
