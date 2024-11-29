package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/controller"
	"github.com/mohammadzaidhussain/bookshop/middleware"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

func RegisterStaffRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	staffService := service.GetStaffService(*repository)
	responseService := utils.GetResponseService()
	staffController := controller.GetStaffController(staffService, *responseService)

	router.POST(
		"/create",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
		}),
		staffController.CreateStaff,
	)

	router.POST(
		"/getAllStaffs",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
		}),
		staffController.GetAllStaffs,
	)

	router.GET(
		"/:staffId",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
			"staff",
		}),
		staffController.GetStaffById,
	)

	router.POST(
		"/:staffId/update",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
			"staff",
		}),
		staffController.UpdateStaffById,
	)

	router.POST(
		"/:staffId/delete",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
		}),
		staffController.DeleteStaffById,
	)
}
