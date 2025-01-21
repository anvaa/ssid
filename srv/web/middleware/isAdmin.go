package middleware

import (
	"github.com/gin-gonic/gin"

	"usr"
)

func IsAdmin(c *gin.Context) {

	user := c.MustGet("user").(users.Users)
	if user.Role != "admin" {
		OnErr(c)
		return
	}
	c.Next()
}
