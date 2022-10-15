package middleware

import (
	authuser "lucy/cashier/lib/auth_user"
	"lucy/cashier/lib/http_response"
	jwthandler "lucy/cashier/lib/jwt_handler"
	"lucy/cashier/lib/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "

	authHeader := c.GetHeader("Authorization")

	if len(authHeader) < len(BEARER_SCHEMA) {
		http_response.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		c.Abort()
		return
	}

	if authHeader[:len(BEARER_SCHEMA)] != BEARER_SCHEMA {
		http_response.JSON(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		c.Abort()
		return
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	claims, err := jwthandler.ValidateToken(tokenString)
	if err != nil {
		http_response.JSON(c, http.StatusUnauthorized, err.Error(), nil)
		c.Abort()
		return
	}

	user, code, err := authuser.FindUser(claims.UserUUID)
	if err != nil {
		logger.Log(logrus.Fields{
			"claims":    claims,
			"user_uuid": claims.UserUUID,
			"user":      user,
		}).Error(err)
		http_response.JSON(c, code, err.Error(), nil)
		c.Abort()
		return
	}

	c.Set("access_token", tokenString)
	c.Set("user_uuid", user.UUID)
	c.Set("branch_uuid", user.BranchUUID)
	c.Set("user_role", user.Role)

	c.Next()
}
