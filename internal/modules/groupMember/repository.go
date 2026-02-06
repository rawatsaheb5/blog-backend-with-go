package groupMember

import (
	"gorm.io/gorm"
)

type Repository interface {
	ListByGroupID(groupID uint64) ([]GroupMember, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) ListByGroupID(groupID uint64) ([]GroupMember, error) {
	var members []GroupMember
	if err := r.db.Where("group_id = ?", groupID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}
