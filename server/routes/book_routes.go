package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/controller"
	"github.com/mohammadzaidhussain/bookshop/middleware"
	repo "github.com/mohammadzaidhussain/bookshop/repository"
	"github.com/mohammadzaidhussain/bookshop/service"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

func RegisterBookRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	bookService := service.GetBookService(*repository)
	responseService := utils.GetResponseService()
	bookController := controller.GetBookController(bookService, *responseService)

	router.POST(
		"/create",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
		}),
		bookController.CreateBook,
	)

	router.POST(
		"/getAllBooks",
		middleware.AuthenticateRequest(false),
		bookController.GetAllBooks,
	)

	router.GET(
		"/:bookId",
		middleware.AuthenticateRequest(false),
		bookController.GetBookById,
	)

	router.POST(
		"/:bookId/update",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
		}),
		bookController.UpdateBookById,
	)

	router.POST(
		"/:bookId/delete",
		middleware.AuthenticateRequest(false),
		middleware.AuthenticateRole([]string{
			"manager",
			"owner",
		}),
		bookController.DeleteBookById,
	)
}
