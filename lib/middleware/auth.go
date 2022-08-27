package middleware

import (
	"fmt"
	"lucy/cashier/lib/http_response"
	jwthandler "lucy/cashier/lib/jwt_handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		http_response.ReturnResponse(c, http.StatusUnauthorized, http.StatusText(401), nil)
		c.Abort()
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	fmt.Println(tokenString)

	claims, err := jwthandler.ValidateToken(tokenString)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}

	fmt.Println(claims)

	c.Set("access_token", tokenString)
	c.Set("user_uuid", claims.UserUUID)
	c.Next()
}
