
type Service struct{
	repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo}
}

func (s *Service) Register(email, password string) error {
    hash := HashPassword(password)
    user := &User{Email: email, Password: hash}
    return s.repo.Create(user)
}