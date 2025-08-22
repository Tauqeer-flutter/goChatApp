package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"goChatApp/handler/responses"
	"net/http"
	"os"
	"strings"
	"time"
)

type JWTClaims struct {
	UserId  string    `json:"user_id"`
	Email   string    `json:"email"`
	Expiry  time.Time `json:"expiry"`
	Created time.Time `json:"created"`
	jwt.RegisteredClaims
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
	tokenString := strings.Trim(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}
	if time.Now().After(claims.Expiry) {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Token expired")
		c.Abort()
		return
	}
	c.Set("user_id", claims.UserId)
	c.Set("email", claims.Email)
	c.Next()
}
