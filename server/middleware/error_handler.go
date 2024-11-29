package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/config"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors[0]
			errMap := make(map[string]any, 0)
			if err != nil {
				splittedErr := strings.Split(err.Error(), "::")
				isDebugEnabled := config.GetEnvProperty("debug")
				if isDebugEnabled != "" && len(splittedErr) > 0 {
					fmt.Printf("Method Name : %s, Status Code : %s, Client Message : %s, Error : %v", splittedErr[2], splittedErr[0], splittedErr[1], splittedErr[3])
				}

				statusCode, err := strconv.Atoi(splittedErr[0])
				if err != nil {
					statusCode = 500
				}
				errMap["message"] = splittedErr[1]
				errMap["statusCode"] = statusCode
				ctx.JSON(statusCode, errMap)
			}
		}
	}
}
