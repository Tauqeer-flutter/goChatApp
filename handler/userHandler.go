package handler

import (
	"goChatApp/domain"
	"goChatApp/handler/requests/auth"
	"goChatApp/handler/responses"
	"goChatApp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserServiceInterface
}

func SetupUserRoutes(router *gin.RouterGroup, userService domain.UserServiceInterface) {
	handler := NewUserHandler(userService)
	group := router.Group("/users/")
	{
		group.POST("/signup", handler.SignUp)
		group.POST("/login", handler.Login)
		group.GET("/list", middlewares.AuthMiddleware, handler.List)
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var request auth.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		//c.JSON(http.StatusBadRequest, responses.BaseResponse{
		//	Success: false,
		//	Message: err.Error(),
		//})
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user := domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		PhotoUrl:  nil,
		Phone:     request.Phone,
	}
	err := h.userService.SignUp(&user)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := h.userService.GenerateJWT(&user)
	if err != nil {
		//c.JSON(500, responses.BaseResponse{
		//	Success: false,
		//	Message: err.Error(),
		//})
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = ""
	responses.SuccessResponse(c, http.StatusOK, "User created successfully", responses.AuthResponse{
		User:        user,
		AccessToken: token,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var request auth.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		//c.JSON(http.StatusBadRequest, responses.BaseResponse{
		//	Success: false,
		//	Message: err.Error(),
		//})
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.userService.Login(request.Email, request.Password)
	if err != nil {
		//c.JSON(http.StatusUnauthorized, responses.BaseResponse{
		//	Success: false,
		//	Message: err.Error(),
		//})
		responses.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	token, err := h.userService.GenerateJWT(user)
	if err != nil {
		//c.JSON(500, responses.BaseResponse{
		//	Success: false,
		//	Message: err.Error(),
		//})
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = ""
	//c.JSON(200, responses.SuccessAuthResponse{
	//	Success:     true,
	//	Message:     "Login successful",
	//	AccessToken: token,
	//	User:        *user,
	//})
	responses.SuccessResponse(c, http.StatusOK, "Login Successful", responses.AuthResponse{
		User:        *user,
		AccessToken: token,
	})
}

func (h *UserHandler) List(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		responses.ErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}
	users, err := h.userService.List()
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	filteredUsers := make([]*responses.UserResponse, 0)
	for _, user := range users {
		if user.Id != userId {
			newUser := &responses.UserResponse{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				PhotoUrl:  user.PhotoUrl,
				Phone:     user.Phone,
			}
			filteredUsers = append(filteredUsers, newUser)
		}
	}
	responses.SuccessResponse(c, http.StatusOK, "Fetched users!", filteredUsers)
}

func NewUserHandler(userService domain.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}
