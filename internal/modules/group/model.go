package group

import "time"

type Group struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"group_id"`
	Title       string    `gorm:"type:text;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	GroupIcon   string    `gorm:"type:text" json:"group_icon,omitempty"`

	AuthorID uint64 `gorm:"not null" json:"author_id"`
	Author   User   `gorm:"foreignKey:AuthorID" json:"author"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
