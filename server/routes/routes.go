package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/middleware"
)

func RegisterRoutes(r *gin.Engine) {

	router := r.Group("/")

	// register error handler middleware
	router.Use(middleware.ErrorHandler())

	userRoutes := router.Group("/users")
	{
		RegisterUserRoutes(userRoutes)
	}

	authRoutes := router.Group("/auth")
	{
		RegisterAuthRoutes(authRoutes)
	}

	bookRoutes := router.Group("/books")
	{
		RegisterBookRoutes(bookRoutes)
	}

	staffRoutes := router.Group("/staffs")
	{
		RegisterStaffRoutes(staffRoutes)
	}

}
