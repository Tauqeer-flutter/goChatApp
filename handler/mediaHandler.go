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
	routerGroup := router.Group("/media", middlewares.AuthMiddleware)
	{
		routerGroup.POST("/upload", handler.Upload)
	}
}

func (ch *MediaHandler) Upload(c *gin.Context) {
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

func NewMediaHandler(service *domain.MediaServiceInterface) *MediaHandler {
	return &MediaHandler{
		service: service,
	}
}
