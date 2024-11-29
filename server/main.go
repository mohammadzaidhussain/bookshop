package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/config"
	"github.com/mohammadzaidhussain/bookshop/middleware"
	"github.com/mohammadzaidhussain/bookshop/routes"
)

func main() {

	app := gin.Default()

	app.Use(gin.Recovery())
	app.Use(middleware.CorsMiddleware)

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":    "Server is running fine!",
			"statusCode": 200,
		})
	})

	routes.RegisterRoutes(app)

	fmt.Println("Go App started successfully !")

	port := config.GetEnvProperty("port")

	fmt.Println(port)

	app.Run(fmt.Sprintf(":%s", port))
}
