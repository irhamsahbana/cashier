package middleware

import (
	"fmt"
	authuser "lucy/cashier/lib/auth_user"
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

	user, code, err := authuser.FindUser(claims.UserUUID)
	if err != nil {
		fmt.Println("masuk sini gan xxx")
		http_response.ReturnResponse(c, code, err.Error(), nil)
		c.Abort()
		return
	}

	c.Set("access_token", tokenString)
	c.Set("user_uuid", user.UUID)
	c.Set("branch_uuid", user.BranchUUID)
	c.Set("user_role", user.Role)

	c.Next()
}
