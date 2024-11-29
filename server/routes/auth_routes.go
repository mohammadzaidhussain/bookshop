package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/controller"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUserService(*repository)
	staffService := service.GetStaffService(*repository)
	authService := service.GetAuthService(*repository, userService, staffService)
	responseService := utils.GetResponseService()
	authController := controller.GetAuthController(authService, *responseService)

	router.POST(
		"/login",
		authController.Login,
	)

	router.POST(
		"/login/staff",
		authController.LoginStaff,
	)

}
