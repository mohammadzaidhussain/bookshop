package utils

import "github.com/gin-gonic/gin"

type ResponseService struct {
	data interface{}
}

/**
	in case of delete pass statusCode as 0, but internally it be 200 and then it will just add one key operation = "delete"
**/
func (rs *ResponseService) Success(ctx *gin.Context, statusCode int, data interface{}, message string) {
	var finalResponse = map[string]any{
		"data":       data,
		"statusCode": statusCode,
		"message":    message,
	}
	if message == "" {
		finalResponse["message"] = "success"
	}
	if statusCode == 0 {
		finalResponse["statusCode"] = 200
		finalResponse["operation"] = "delete"
	}

	ctx.JSON(statusCode, finalResponse)
}

func GetResponseService() *ResponseService {
	return &ResponseService{}
}
