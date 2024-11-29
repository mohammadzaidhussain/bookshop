package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/model"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

type UserController struct {
	userService     service.IUserService
	responseService utils.ResponseService
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var dto model.UserCreateDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "UserController_CreateUser", err))
		return
	}

	data, err := uc.userService.CreateUser(dto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	uc.responseService.Success(ctx, 200, data, "sucessfully saved!")

}

func GetUserController(userService service.IUserService, responseService utils.ResponseService) *UserController {
	return &UserController{
		userService:     userService,
		responseService: responseService,
	}
}
