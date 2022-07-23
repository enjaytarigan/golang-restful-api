package auth

import (
	"brodo-demo/api/common"
	"brodo-demo/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *auth.AuthService
}

func NewAuthController(service *auth.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

type PostLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostLoginResponse struct {
	AccessToken string `json:"accessToken"`
}

func (c *AuthController) PostLogin(ctx *gin.Context) {
	payload := PostLoginPayload{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	token, err := c.authService.Login(payload.Username, payload.Password)

	if err != nil {
		code, body := common.NewErrorServiceResponse(err)
		ctx.JSON(code, body)
		return
	}

	body := PostLoginResponse{
		AccessToken: token,
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse(body))
}