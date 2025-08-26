package handler

import (
	"goChatApp/domain"
	"goChatApp/handler/requests"
	"goChatApp/handler/responses"
	"goChatApp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatService *domain.ChatServiceInterface
}

func SetupChatRoutes(router *gin.RouterGroup, service *domain.ChatServiceInterface) {
	handler := NewChatHandler(service)
	routerGroup := router.Group("/chats", middlewares.AuthMiddleware)
	{
		routerGroup.POST("/send-message", handler.SendMessage)
	}
}

func (ch ChatHandler) SendMessage(c *gin.Context) {
	var request requests.SendMessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Bad Request")
		return
	}
	err := (*ch.chatService).SendMessage(&request)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.SuccessResponse(c, http.StatusOK, "Message sent!", nil)
}

func NewChatHandler(service *domain.ChatServiceInterface) *ChatHandler {
	return &ChatHandler{
		chatService: service,
	}
}
