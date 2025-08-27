package handler

import (
	"fmt"
	"goChatApp/domain"
	"goChatApp/handler/requests/group"
	"goChatApp/handler/responses"
	"goChatApp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
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
		groups.GET("/list", handler.List)
	}
}

func (h GroupHandler) Create(c *gin.Context) {
	var request group.CreateGroupRequest
	currentUserId, exists := c.Get("user_id")
	if !exists {
		responses.ErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err)
		responses.ErrorResponse(c, http.StatusBadRequest, "Invalid request body!")
		return
	}
	var memberCount int
	if request.GroupType == "private" {
		memberCount = 2
		if request.OtherUserId == nil {
			responses.ErrorResponse(c, http.StatusBadRequest, "other_user_id is required for private group!")
			return
		}
	} else {
		memberCount = 0
	}
	group := domain.Group{
		Name:        request.Name,
		Description: request.Description,
		GroupType:   request.GroupType,
		MemberCount: memberCount,
	}
	userId := currentUserId.(int64)
	createdGroup, err := h.groupService.Create(&group, &userId, request.OtherUserId)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.SuccessResponse(c, http.StatusOK, "Group created!", createdGroup)
}

func (h GroupHandler) List(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		responses.ErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}
	groups, err := h.groupService.List(userId.(int64))
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.SuccessResponse(c, http.StatusOK, "Fetched groups!", groups)
}
