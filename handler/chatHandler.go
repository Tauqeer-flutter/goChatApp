package handler

import (
	"fmt"
	"goChatApp/domain"
	"goChatApp/handler/requests/chat"
	"goChatApp/handler/responses"
	"goChatApp/middlewares"
	"goChatApp/services"
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
		//routerGroup.POST("/send-message", handler.SendMessage)
		routerGroup.GET("/list", handler.List)
		routerGroup.GET("/ws", handler.ChatWS)
	}
}

//func (ch ChatHandler) SendMessage(c *gin.Context) {
//	var request requests.SendMessageRequest
//	if err := c.ShouldBindJSON(&request); err != nil {
//		responses.ErrorResponse(c, http.StatusBadRequest, "Bad Request")
//		return
//	}
//
//	err := (*ch.chatService).SendMessage(&request)
//	if err != nil {
//		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	responses.SuccessResponse(c, http.StatusOK, "Message sent!", nil)
//}

func (ch ChatHandler) List(c *gin.Context) {
	var request requests.AllMessagesRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Group ID is required")
		return
	}
	chats, err := (*ch.chatService).List(request.GroupId)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.SuccessResponse(c, http.StatusOK, "Fetched chats!", chats)
}

// ChatWS handles WebSocket connections for chat events
func (ch ChatHandler) ChatWS(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		responses.ErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}
	var request requests.ChatWebsocketRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	service := (*ch.chatService).(*services.ChatService)
	_, err := service.GroupRepository.GetGroupById(&request.GroupId)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Group not found")
		return
	}
	ws, err := domain.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer ws.Close()
	domain.Mutex.Lock()
	client := domain.Client{
		Conn:   ws,
		Active: true,
		UserId: userId.(int64),
	}
	domain.Clients[&client] = request.GroupId
	domain.Mutex.Unlock()
	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			domain.Mutex.Lock()
			delete(domain.Clients, &client)
			domain.Mutex.Unlock()
			break
		}
		message := string(bytes)
		fmt.Println(message)
		messageRequest := requests.SendMessageRequest{
			SenderId: userId.(int64),
			GroupId:  &request.GroupId,
			Message:  message,
		}
		newChat, err := (*ch.chatService).SendMessage(&messageRequest)
		if err != nil {
			err := ws.WriteJSON(responses.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			if err != nil {
				domain.Mutex.Lock()
				delete(domain.Clients, &client)
				domain.Mutex.Unlock()
				break
			}
			continue
		}
		for allClient, groupId := range domain.Clients {
			if groupId == request.GroupId {
				err := allClient.Conn.WriteJSON(responses.Response{
					Status:  true,
					Message: "New message received",
					Data:    newChat,
				})
				if err != nil {
					domain.Mutex.Lock()
					delete(domain.Clients, allClient)
					domain.Mutex.Unlock()
					break
				}
			}
		}
	}
}

func NewChatHandler(service *domain.ChatServiceInterface) *ChatHandler {
	return &ChatHandler{
		chatService: service,
	}
}
