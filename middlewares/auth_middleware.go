package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"goChatApp/handler/responses"
	"net/http"
	"os"
	"strings"
	"time"
)

type JWTClaims struct {
	UserId  int64  `json:"user_id"`
	Email   string `json:"email"`
	Expiry  int64  `json:"expiry"`
	Created int64  `json:"created"`
	jwt.RegisteredClaims
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println(err)
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
	expiry := time.Unix(claims.Expiry, 0)
	if time.Now().After(expiry) {
		responses.ErrorResponse(c, http.StatusUnauthorized, "Token expired")
		c.Abort()
		return
	}
	c.Set("user_id", claims.UserId)
	c.Set("email", claims.Email)
	c.Next()
}
