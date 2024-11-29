package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/dto"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

type AuthController struct {
	authService     service.IAuthService
	responseService utils.ResponseService
}

func (ac *AuthController) Login(ctx *gin.Context) {

	var loginDto dto.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "AuthController_Login", err))
	}

	data, err := ac.authService.Login(loginDto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	ac.responseService.Success(ctx, 200, data, "sucessfully saved!")
}

func (ac *AuthController) LoginStaff(ctx *gin.Context) {

	var loginDto dto.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "AuthController_LoginStaff", err))
	}

	data, err := ac.authService.LoginStaff(loginDto, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	ac.responseService.Success(ctx, 200, data, "sucessfully saved!")
}

func GetAuthController(authService service.IAuthService, responseService utils.ResponseService) *AuthController {
	return &AuthController{
		authService:     authService,
		responseService: responseService,
	}
}
