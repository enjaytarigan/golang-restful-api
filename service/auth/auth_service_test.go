package auth

import (
	// security "brodo-demo/service/security/mocks"
	"brodo-demo/entity"
	repository "brodo-demo/repository/mocks"
	security "brodo-demo/service/security/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_UsernameNotRegistered(t *testing.T) {
	t.Run("should return ErrInvalidCredential when given not registered username", func(t *testing.T) {
		var mockUserRepository = new(repository.UserRepository)

		username := "adminbrodo"

		mockUserRepository.On("FindByUsername", username).Return(entity.User{}, errors.New("username not found"))

		authSerivce := AuthService{
			userRepository: mockUserRepository,
		}

		_, err := authSerivce.Login(username, "supersecretpassword")

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrInvalidCredential)
	})
}

func TestLogin_WrongPassword(t *testing.T) {
	t.Run("should return ErrInvalidCredential when given wrong password", func(t *testing.T) {
		userRepository := new(repository.UserRepository)
		passwordHash := new(security.PasswordHash)

		service := AuthService{
			userRepository: userRepository,
			passwordHash:   passwordHash,
		}

		user := entity.User{
			Username: "randomusername",
			Password: "supersecretpassword",
		}

		userRepository.On("FindByUsername", mock.AnythingOfType("string")).Return(user, nil)

		passwordHash.On("IsMatch", user.Password, mock.AnythingOfType("string")).Return(false)

		_, err := service.Login(user.Username, user.Password)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrInvalidCredential)
	})
}

func TestLogin_Success(t *testing.T) {
	t.Run("should return access token when given valid credential", func(t *testing.T) {
		mockUserRepository := new(repository.UserRepository)
		mockPasswordHash := new(security.PasswordHash)
		mockTokenManager := new(security.AuthenticationTokenManager)

		payload := struct {
			username string
			password string
		}{
			username: "brododemotest",
			password: "passwordtestbrodo",
		}

		user := entity.User{
			ID:       1,
			Username: "brododemotest",
			Password: "encrypted_password",
		}

		mockUserRepository.On("FindByUsername", payload.username).Return(user, nil)
		mockPasswordHash.On("IsMatch", payload.password, user.Password).Return(true)
		mockTokenManager.On("CreateAccessToken", user.ID).Return("access_token", nil)

		service := AuthService{
			userRepository: mockUserRepository,
			passwordHash:   mockPasswordHash,
			tokenManager:   mockTokenManager,
		}

		accessToken, err := service.Login(payload.username, payload.password)

		assert.Nil(t, err)
		assert.NotEmpty(t, accessToken)
	})
}
