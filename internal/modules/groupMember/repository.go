package groupMember

import (
	"gorm.io/gorm"
)

type Repository interface {
	ListByGroupID(groupID uint64) ([]GroupMember, error)
	ListGroupIDsByUserID(userID uint64) ([]uint64, error)
	UpdateStatus(groupID uint64, userID uint64, status string) (int64, error)
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

func (r *repository) ListGroupIDsByUserID(userID uint64) ([]uint64, error) {
	var groupIDs []uint64
	rows := []GroupMember{}
	if err := r.db.Select("group_id").Where("user_id = ?", userID).Find(&rows).Error; err != nil {
		return nil, err
	}
	groupIDs = make([]uint64, 0, len(rows))
	for _, gm := range rows {
		groupIDs = append(groupIDs, gm.GroupID)
	}
	return groupIDs, nil
}

func (r *repository) UpdateStatus(groupID uint64, userID uint64, status string) (int64, error) {
	res := r.db.Model(&GroupMember{}).Where("group_id = ? AND user_id = ?", groupID, userID).Update("status", status)
	return res.RowsAffected, res.Error
}
