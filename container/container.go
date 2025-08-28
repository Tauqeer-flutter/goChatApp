package container

import (
	"goChatApp/config"
	"goChatApp/domain"
	"goChatApp/handler/routes"
	"goChatApp/repositories"
	"goChatApp/services"

	"github.com/gin-gonic/gin"
)

type Container struct {
	Config       *config.Config
	UserService  domain.UserServiceInterface
	GroupService domain.GroupServiceInterface
	ChatService  domain.ChatServiceInterface
	MediaService domain.MediaServiceInterface
}

func (c *Container) SetupRoutes(router *gin.Engine) {
	group := router.Group("/api")
	{
		routes.SetupUserRoutes(group, c.UserService)
		routes.SetupGroupRoutes(group, c.GroupService)
		routes.SetupChatRoutes(group, &c.ChatService)
		routes.SetupMediaRoutes(group, &c.MediaService)
	}
}

func NewContainer() *Container {
	cfg := config.LoadConfig()
	userRepository := repositories.NewUserRepository(cfg.DB)
	userService := services.NewUserService(&userRepository)
	groupRepository := repositories.NewGroupRepository(cfg.DB)
	groupService := services.NewGroupService(&groupRepository, &userRepository)
	chatRepository := repositories.NewChatRepository(cfg.DB)
	chatService := services.NewChatService(&chatRepository, &groupRepository, &userRepository)
	mediaService := services.NewMediaService()
	return &Container{
		Config:       cfg,
		UserService:  userService,
		GroupService: groupService,
		ChatService:  chatService,
		MediaService: mediaService,
	}
}
