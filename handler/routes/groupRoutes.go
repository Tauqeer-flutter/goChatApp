package routes

import (
	"goChatApp/domain"
	"goChatApp/handler"

	"github.com/gin-gonic/gin"
)

func SetupGroupRoutes(group *gin.RouterGroup, service domain.GroupServiceInterface) {
	handler.SetupGroupRoutes(group, service)
}
