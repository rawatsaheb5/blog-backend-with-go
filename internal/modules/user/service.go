package user


import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
	jwtSecret string
}

func NewService(repo Repository, jwtSecret string) *Service {
	return &Service{
		repo: repo,
		jwtSecret: jwtSecret,
		}
}

func (s *Service) Register(email, password string) error {
	hash := HashPassword(password)
	user := &User{Email: email, Password: hash}
	return s.repo.CreateUser(user)
}

func (s *Service) Login(email, password string) (*User, string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	if !VerifyPassword(password, user.Password) {
		return nil, "", errors.New("invalid password")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil , "", errors.New("failed to generate token")
	}
	return user, token, nil
}

// HashPassword - placeholder function, implement proper hashing
func HashPassword(password string) string {
	// TODO: Implement proper password hashing (e.g., bcrypt)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return password
	}
	return string(bytes)
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) generateToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
