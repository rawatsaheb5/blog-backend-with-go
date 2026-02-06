package group

type Service interface {
	CreateGroup(title string, authorID uint64) (*Group, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateGroup(title string, authorID uint64) (*Group, error) {
	g := &Group{
		Title:    title,
		AuthorID: authorID,
	}
	if err := s.repo.Create(g); err != nil {
		return nil, err
	}
	return g, nil
}
