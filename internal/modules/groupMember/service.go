package groupMember

type Service interface {
	GetAllGroupMembers(groupID uint64) ([]GroupMember, error)
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
