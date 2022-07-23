package user

import (
	"brodo-demo/entity"
	"brodo-demo/repository"
	"brodo-demo/service/security"
	"errors"
	"strings"
)

var (
	ErrInvalidUsername     = errors.New("invalid username")
	ErrPasswordTooShort    = errors.New("password length too short (minimum 8 character)")
	ErrUsernameAlreadyUsed = errors.New("username already used")
	ErrUnexpected          = errors.New("error create user")
)

type UserService struct {
	userRepo     repository.UserRepository
	passwordHash security.PasswordHash
}

func NewUserService(userRepository repository.UserRepository, passwordHash security.PasswordHash) UserService {
	return UserService{
		userRepo:     userRepository,
		passwordHash: passwordHash,
	}
}

type CreateUserPayload struct {
	Username string
	Password string
}

func (service *UserService) CreateUser(payload CreateUserPayload) (int, error) {
	var blankUserId int

	lowerUsername := strings.ToLower(payload.Username)

	if len(lowerUsername) < 5 {
		return blankUserId, ErrInvalidUsername
	}

	if len(payload.Password) < 8 {
		return blankUserId, ErrPasswordTooShort
	}

	isAvailable := service.userRepo.VerifyAvailableUsername(lowerUsername)

	if !isAvailable {
		return blankUserId, ErrUsernameAlreadyUsed
	}

	hashedPassword, err := service.passwordHash.Hash(payload.Password)

	if err != nil {
		return blankUserId, ErrUnexpected
	}

	user := entity.User{
		Username: lowerUsername,
		Password: hashedPassword,
	}

	userId, err := service.userRepo.Insert(user)

	if err != nil {
		return blankUserId, ErrUnexpected
	}

	return userId, nil
}