package groupMember
import "time"


type GroupMember struct {
	// Composite Primary Key
	GroupID uint64 `gorm:"primaryKey" json:"group_id"`
	UserID  uint64 `gorm:"primaryKey" json:"user_id"`

	// Member status
	Status string `gorm:"type:varchar(20);not null;default:'active'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
