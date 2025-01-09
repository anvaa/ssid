package middleware

import (
	"github.com/gin-gonic/gin"

	"users"
)

func IsAdmin(c *gin.Context) {

	user := c.MustGet("user").(users.Users)
	if user.Role != "admin" {
		OnErr(c)
		return
	}
	c.Next()
}
