package routes

import (
	"github.com/gin-gonic/gin"
	"goChatApp/domain"
	"goChatApp/handler"
)

func SetupGroupRoutes(group *gin.RouterGroup, service domain.GroupServiceInterface) {
	handler.SetupGroupRoutes(group, service)
}
