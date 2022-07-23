package user

import (
	"brodo-demo/api/common"
	"brodo-demo/api/user/request"
	"brodo-demo/api/user/response"
	"brodo-demo/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *user.UserService
}

func NewUserContoller(userService *user.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) PostUser(ctx *gin.Context) {
	body := request.PostUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	userId, err := controller.userService.CreateUser(body.ToCreateUserPayload())

	if err != nil {
		statusCode, body := common.NewErrorServiceResponse(err)
		ctx.JSON(statusCode, body)
		return
	}

	response := response.PostUserResponse{
		UserId: userId,
	}

	ctx.JSON(http.StatusCreated, common.NewSuccessResponse(response))
}