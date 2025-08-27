package routes

import (
	"goChatApp/domain"
	"goChatApp/handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(group *gin.RouterGroup, userService domain.UserServiceInterface) {
	group.Any("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "message": "Welcome to Go Chat APIs"})
	})
	handler.SetupUserRoutes(group, userService)
}
