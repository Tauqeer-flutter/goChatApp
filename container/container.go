package container

import (
	"github.com/gin-gonic/gin"
	"goChatApp/config"
	"goChatApp/domain"
	"goChatApp/handler/routes"
	"goChatApp/repositories"
	"goChatApp/services"
)

type Container struct {
	Config         *config.Config
	UserService    domain.UserServiceInterface
	UserRepository domain.UserRepositoryInterface
	GroupService   domain.GroupServiceInterface
}

func (c *Container) SetupRoutes(router *gin.Engine) {
	group := router.Group("/api")
	{
		routes.SetupUserRoutes(group, c.UserService)
		routes.SetupGroupRoutes(group, c.GroupService)
	}
}

func NewContainer() *Container {
	cfg := config.LoadConfig()
	userRepository := repositories.NewUserRepository(cfg.DB)
	userService := services.NewUserService(userRepository)
	groupRepository := repositories.NewGroupRepository(cfg.DB)
	groupService := services.NewGroupService(groupRepository, userRepository)
	return &Container{
		Config:       cfg,
		UserService:  userService,
		GroupService: groupService,
	}
}
