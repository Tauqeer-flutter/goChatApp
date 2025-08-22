package handler

import (
	"github.com/gin-gonic/gin"
	"goChatApp/domain"
	"goChatApp/handler/requests"
	"goChatApp/handler/responses"
	"goChatApp/middlewares"
	"net/http"
)

type GroupHandler struct {
	groupService domain.GroupServiceInterface
}

func NewGroupHandler(service domain.GroupServiceInterface) *GroupHandler {
	return &GroupHandler{
		groupService: service,
	}
}

func SetupGroupRoutes(e *gin.RouterGroup, service domain.GroupServiceInterface) {
	handler := NewGroupHandler(service)

	groups := e.Group("/groups", middlewares.AuthMiddleware)
	{
		groups.POST("/create", handler.Create)
	}
}

func (h GroupHandler) Create(c *gin.Context) {
	var request requests.CreateGroupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if request.GroupType == "private" {
		if request.OtherUserId == "" {
			responses.ErrorResponse(c, http.StatusBadRequest, "other_user_id is required for private group!")
			return
		}
	}
	group := domain.Group{
		Name:        request.Name,
		Description: request.Description,
		GroupType:   request.GroupType,
	}
	createdGroup, err := h.groupService.Create(&group)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.SuccessResponse(c, http.StatusOK, "Group created!", createdGroup)
}
