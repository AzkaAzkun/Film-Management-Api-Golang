package controller

import (
	"film-management-api-golang/internal/api/service"
	"film-management-api-golang/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		GetById(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUser(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetById(ctx *gin.Context) {
	userId := ctx.Param("id")
	requesterId, _ := ctx.Get("user_id")
	requesterIdStr, _ := requesterId.(string)

	result, err := c.userService.GetById(ctx.Request.Context(), userId, requesterIdStr)
	if err != nil {
		response.NewFailed("failed get detail user", err).Send(ctx)
		return
	}

	response.NewSuccess("success get detail user", result).Send(ctx)
}
