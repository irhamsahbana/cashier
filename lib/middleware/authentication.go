package middleware

import (
	"lucy/cashier/lib/http_response"
	jwthandler "lucy/cashier/lib/jwt_handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || authHeader[:len(BEARER_SCHEMA)] != BEARER_SCHEMA {
		http_response.ReturnResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		c.Abort()
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	claims, err := jwthandler.ValidateToken(tokenString)
	if err != nil {
		http_response.ReturnResponse(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}

	c.Set("access_token", tokenString)
	c.Set("user_uuid", claims.UserUUID)
	c.Set("user_role", claims.Role)

	c.Next()
}
