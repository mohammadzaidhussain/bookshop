package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/controller"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUserService(*repository)
	responseService := utils.GetResponseService()
	userController := controller.GetUserController(userService, *responseService)

	router.POST(
		"/add",
		userController.CreateUser,
	)
}
