package routes

import (
	"goChatApp/domain"
	"goChatApp/handler"

	"github.com/gin-gonic/gin"
)

func SetupChatRoutes(c *gin.RouterGroup, service *domain.ChatServiceInterface) {
	handler.SetupChatRoutes(c, service)
}
