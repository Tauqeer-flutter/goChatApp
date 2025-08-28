package routes

import (
	"goChatApp/domain"
	"goChatApp/handler"

	"github.com/gin-gonic/gin"
)

func SetupMediaRoutes(c *gin.RouterGroup, service *domain.MediaServiceInterface) {
	handler.SetupMediaRoutes(c, service)
}
