package auth

import (
	"brodo-demo/repository"
	"brodo-demo/service/security"
	"errors"
	"strings"
)

var (
	ErrInvalidCredential = errors.New("username or password is wrong")
)

type AuthService struct {
	userRepository repository.UserRepository
	passwordHash   security.PasswordHash
	tokenManager   security.AuthenticationTokenManager
}

func NewAuthService(userRepo repository.UserRepository, passwordHash security.PasswordHash, tokenManager security.AuthenticationTokenManager) *AuthService {
	return &AuthService{
		userRepository: userRepo,
		passwordHash:   passwordHash,
		tokenManager:   tokenManager,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepository.FindByUsername(strings.ToLower(username))

	if err != nil {
		return "", ErrInvalidCredential
	}

	if !s.passwordHash.IsMatch(password, user.Password) {
		return "", ErrInvalidCredential
	}

	token, err := s.tokenManager.CreateAccessToken(user.ID)

	if err != nil {

		return token, err
	}

	return token, nil
}
