package groupMember

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Service interface {
	GetAllGroupMembers(groupID uint64) ([]GroupMember, error)
	GetUserGroupIDs(userID uint64) ([]uint64, error)
	LeaveGroup(groupID uint64, userID uint64) (bool, error)
	GenerateInviteLink(groupID uint64, inviterID uint64, inviteeEmail string) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllGroupMembers(groupID uint64) ([]GroupMember, error) {
	return s.repo.ListByGroupID(groupID)
}

func (s *service) GetUserGroupIDs(userID uint64) ([]uint64, error) {
	return s.repo.ListGroupIDsByUserID(userID)
}

func (s *service) LeaveGroup(groupID uint64, userID uint64) (bool, error) {
	affected, err := s.repo.UpdateStatus(groupID, userID, "LEFT")
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// GenerateInviteLink creates a simple opaque token-based invite link. In production, persist tokens and send emails.
func (s *service) GenerateInviteLink(groupID uint64, inviterID uint64, inviteeEmail string) (string, error) {
	// Simple, non-persisted token generation for demonstration
	payload := fmt.Sprintf("%d:%d:%s:%d", groupID, inviterID, inviteeEmail, time.Now().UnixNano())
	sum := sha256.Sum256([]byte(payload))
	token := hex.EncodeToString(sum[:])
	link := fmt.Sprintf("/api/group/join?token=%s", token)
	return link, nil
}
