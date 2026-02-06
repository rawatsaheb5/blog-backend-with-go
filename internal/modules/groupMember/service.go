package groupMember

type Service interface {
	GetAllGroupMembers(groupID uint64) ([]GroupMember, error)
	GetUserGroupIDs(userID uint64) ([]uint64, error)
	LeaveGroup(groupID uint64, userID uint64) (bool, error)
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
