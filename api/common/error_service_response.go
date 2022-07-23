package common

import (
	"brodo-demo/service/auth"
	"brodo-demo/service/user"
	"net/http"
)

type ErrorServiceResponse struct {
	Status  bool   `json:"status"`
	Message string `jsn:"message"`
}

func NewErrorServiceResponse(err error) (statusCode int, body ErrorServiceResponse) {
	return translateError(err)
}

func translateError(err error) (statusCode int, body ErrorServiceResponse) {
	switch err {
	case user.ErrUnexpected:
		return newInternalServerError()

	case user.ErrInvalidUsername:
		return newInvalidSpecError("invalid username")

	case user.ErrUsernameAlreadyUsed:
		return newInvalidSpecError("username already exists")

	case user.ErrPasswordTooShort:
		return newInvalidSpecError("password too short")

	case auth.ErrInvalidCredential:
		return newInvalidSpecError("username or password is wrong")

	default:
		return newInternalServerError()
	}
}

func newInvalidSpecError(message string) (statusCode int, body ErrorServiceResponse) {
	return http.StatusBadRequest, ErrorServiceResponse{
		Status:  false,
		Message: message,
	}
}

func newInternalServerError() (statusCode int, body ErrorServiceResponse) {
	return http.StatusInternalServerError, ErrorServiceResponse{
		Status:  false,
		Message: "Internal Server Error",
	}
}
