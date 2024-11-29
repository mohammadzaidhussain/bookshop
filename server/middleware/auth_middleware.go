package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mohammadzaidhussain/bookshop/utils"
)

func AuthenticateRequest(bypass bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")

		if token != "" {
			token = strings.Split(token, " ")[1]
		}

		if token == "" || token == "null" || token == "undefined" {
			token = ctx.Request.Header.Get("token")
		}

		if token == "" || token == "null" || token == "undefined" {
			token, _ = ctx.Cookie("token")
		}

		if token == "" {
			token, _ = ctx.GetQuery("token")
		}

		if bypass && token == "" {
			ctx.Next()
			return
		}

		if token == "" {

			ctx.JSON(401, gin.H{
				"message":    "unauthorized",
				"statusCode": 401,
			})

			ctx.Abort()

		} else {
			jwtToken, err := utils.VerifyToken(token)

			if err != nil {
				ctx.JSON(401, gin.H{
					"message":    "unauthorized",
					"statusCode": 401,
				})

				ctx.Abort()

			} else {

				if userId, ok := (*jwtToken)["user_id"]; ok {

					role := (*jwtToken)["role"]

					ctx.Set("role", role)
					ctx.Set("user_id", userId)

					ctx.Next()
					return

				}
			}
		}
	}
}

func AuthenticateRole(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctxRole := ctx.GetString("role")
		var isRoleMatched = false
		for _, role := range roles {
			if role == ctxRole {
				isRoleMatched = true
				break
			}
		}
		if !isRoleMatched {
			ctx.JSON(401, gin.H{
				"message":    "unauthorized",
				"statusCode": 401,
			})

			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
