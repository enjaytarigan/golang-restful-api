package request

import (
	"brodo-demo/service/user"
)

type PostUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (body *PostUserRequestBody) ToCreateUserPayload() user.CreateUserPayload {
	return user.CreateUserPayload{
		Username: body.Username,
		Password: body.Password,
	}
}