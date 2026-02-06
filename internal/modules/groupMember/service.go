package groupMember

type Service interface {
	GetAllGroupMembers(groupID uint64) ([]GroupMember, error)
	GetUserGroupIDs(userID uint64) ([]uint64, error)
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
