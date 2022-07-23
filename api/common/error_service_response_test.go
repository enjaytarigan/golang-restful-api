package common

import (
	"brodo-demo/service/auth"
	"brodo-demo/service/user"
	"errors"
	"net/http"
	"testing"
)

func TestErrorServiceResponse(t *testing.T) {
	testTable := []struct {
		testName         string
		errInput         error
		statusCodeWant int       
		messageWant    string
	}{
		{
			testName:         "should have 400 status code when translate error:ErrInvalidUsername",
			errInput:         user.ErrInvalidUsername,
			statusCodeWant: http.StatusBadRequest,
			messageWant:    "invalid username",
		},
		{
			testName:         "should have 400 status code when translate error:ErrPasswordTooShort",
			errInput:         user.ErrPasswordTooShort,
			statusCodeWant: http.StatusBadRequest,
			messageWant:    "password too short",
		},
		{
			testName:         "should have 400 status code when translate error:ErrUsernameAlreadyUsed",
			errInput:         user.ErrUsernameAlreadyUsed,
			statusCodeWant: http.StatusBadRequest,
			messageWant:    "username already exists",
		},
		{
			testName: "should have 500 status code when translate error: ErrUnexpected",
			errInput: user.ErrUnexpected,
			statusCodeWant: http.StatusInternalServerError,
			messageWant: "Internal Server Error",
		},
		{
			testName: "should have 400 status code when translate error:ErrInvalidCredential",
			errInput: auth.ErrInvalidCredential,
			statusCodeWant: http.StatusBadRequest,
			messageWant: "username or password is wrong",
		},
		{
			testName: "should have 500 status code when given random error",
			errInput: errors.New("Internal Server Error"),
			statusCodeWant: http.StatusInternalServerError,
			messageWant: "Internal Server Error",
		},
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			statusCode, errServiceResponse := NewErrorServiceResponse(test.errInput)

			if statusCode != test.statusCodeWant {
				t.Errorf("NewErrorServiceResponse().statusCode = %d, want = %d", statusCode, test.statusCodeWant)
			}

			if errServiceResponse.Message !=  test.messageWant {
				t.Errorf("NewErrorServiceResponse().Message = %s, want = %s", errServiceResponse.Message, test.errInput.Error())
			}
		})
	}
}