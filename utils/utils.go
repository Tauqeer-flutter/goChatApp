package utils

import (
	"goChatApp/handler/responses"

	"github.com/gin-gonic/gin"
)

func Must[T any](c *gin.Context, statusCode int, err error) {
	if err != nil {
		responses.ErrorResponse(c, statusCode, err.Error())
		return
	}
}
