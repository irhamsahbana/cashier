package middleware

import (
	"lucy/cashier/lib/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRole string

const (
	UserRole_OWNER        UserRole = "Owner"
	UserRole_BRANCH_OWNER UserRole = "Branch Owner"
	UserRole_MANAGER      UserRole = "Manager"

	userRole_HEAD_CASHIER    UserRole = "Admin Cashier"
	UserRole_HEAD_ACCOUNTANT UserRole = "Head Accountant"
	UserRole_STOCK_OVERSIER  UserRole = "Stock Overseer"

	UserRole_CASHIER      UserRole = "Cashier"
	UserRole_ACCOUNTANT   UserRole = "Accountant"
	UserRole_STOCK_KEEPER UserRole = "Stock Keeper"
)

func Authorization(permittedRoles []UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, role := range permittedRoles {
			if role == UserRole(c.GetString("user_role")) {
				c.Next()
				return
			}
		}

		http_response.ReturnResponse(c, http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
		c.Abort()
	}
}
