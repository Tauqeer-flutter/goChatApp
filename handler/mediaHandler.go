package handler

import (
	"goChatApp/domain"
	requests "goChatApp/handler/requests/chat"
	"goChatApp/handler/responses"
	"goChatApp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	service *domain.MediaServiceInterface
}

func SetupMediaRoutes(router *gin.RouterGroup, service *domain.MediaServiceInterface) {
	handler := NewMediaHandler(service)
	routerGroup := router.Group("/media")
	{
		routerGroup.POST("/upload-chat-file", middlewares.AuthMiddleware, handler.UploadChatFile)
		routerGroup.GET("/:groupId/:filename", handler.GetChatFile)
	}
}

func (ch *MediaHandler) UploadChatFile(c *gin.Context) {
	var request requests.GroupIdRequest
	if err := c.ShouldBind(&request); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "File is required")
		return
	}
	filename, err := (*ch.service).UploadChatFile(request.GroupId, file)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.SuccessResponse(c, http.StatusOK, "File uploaded", filename)
}

func (ch *MediaHandler) GetChatFile(c *gin.Context) {
	filename := c.Params.ByName("filename")
	if filename == "" {
		responses.ErrorResponse(c, http.StatusBadRequest, "Filename is required")
		return
	}
	groupId := c.Params.ByName("groupId")
	if groupId == "" {
		responses.ErrorResponse(c, http.StatusBadRequest, "Group ID is required")
		return
	}
	path := "uploads/chats/" + groupId + "/" + filename
	c.File(path)
}

func NewMediaHandler(service *domain.MediaServiceInterface) *MediaHandler {
	return &MediaHandler{
		service: service,
	}
}
