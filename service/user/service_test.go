package user

import (
	"strings"
	"testing"

	"brodo-demo/entity"
	repository "brodo-demo/repository/mocks"
	"brodo-demo/service/security/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_InvalidPayload(t *testing.T) {

	var userService = UserService{
		userRepo:     nil,
		passwordHash: nil,
	}
	invalidPayloadTests := []struct {
		name       string
		payload    CreateUserPayload
		wantUserId int
		wantErr    error
	}{
		{
			name: "should return ErrInvalidUsername when username is empty string",
			payload: CreateUserPayload{
				Username: "",
				Password: "supersecretpassword",
			},
			wantUserId: 0,
			wantErr:    ErrInvalidUsername,
		},
		{
			name: "should return ErrInvalidUsername when length of username is less than 5 ",
			payload: CreateUserPayload{
				Username: "demo",
				Password: "supersecretpassword",
			},
			wantUserId: 0,
			wantErr:    ErrInvalidUsername,
		},
		{
			name: "should return ErrPasswordTooShort when length of password is less than 8",
			payload: CreateUserPayload{
				Username: "brododemo",
				Password: "short",
			},
			wantUserId: 0,
			wantErr:    ErrPasswordTooShort,
		},
	}

	for _, test := range invalidPayloadTests {
		t.Run(test.name, func(t *testing.T) {
			userId, err := userService.CreateUser(test.payload)

			assert.Equal(t, 0, userId)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, test.wantErr)
		})
	}
}

func TestCreateUser_Failed_UsernameIsNotAvailable(t *testing.T) {
	mockUserRepository := repository.UserRepository{}



	payload := CreateUserPayload{
		Username: "alreadyexistusername",
		Password: "passwordstrong",
	}

	mockUserRepository.On("VerifyAvailableUsername", strings.ToLower(payload.Username)).Return(false)

	var userService = UserService{
		userRepo:     &mockUserRepository,
		passwordHash: nil,
	}

	userId, err := userService.CreateUser(payload)

	assert.Empty(t, userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrUsernameAlreadyUsed)
}

func TestCreateUser_Success(t *testing.T) {
	payload := CreateUserPayload{
		Username: "brodouser",
		Password: "supersecretpassword",
	}
	mockUserRepository := new(repository.UserRepository)
	mockPasswordHash := new(mocks.PasswordHash)

	mockUserRepository.On("VerifyAvailableUsername", strings.ToLower(payload.Username)).Return(true)

	mockPasswordHash.On("Hash", payload.Password).Return("encrypted_password", nil)

	mockUserRepository.On("Insert", entity.User{
		Username: strings.ToLower(payload.Username),
		Password: "encrypted_password",
	}).Return(1, nil)

	var userService = UserService{
		userRepo:     mockUserRepository,
		passwordHash: mockPasswordHash,
	}

	userId, err := userService.CreateUser(payload)

	assert.Nil(t, err)
	assert.NotEmpty(t, userId)
}
