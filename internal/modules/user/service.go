package user
type Service struct{
	repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo}
}

func (s *Service) Register(email, password string) error {
    hash := HashPassword(password)
    user := &User{Email: email, Password: hash}
    return s.repo.CreateUser(user)
}

// HashPassword - placeholder function, implement proper hashing
func HashPassword(password string) string {
	// TODO: Implement proper password hashing (e.g., bcrypt)
	return password
}