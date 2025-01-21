package middleware

import (
	"log"
	"usr"

	"github.com/gin-gonic/gin"

	"net/http"
)

func OnErr(c *gin.Context) {

	// error handling here
	log.Println("Error: ", c.Errors.String())

	Logout(c)
}

func Logout(c *gin.Context) {
	// delete cookie from browser, redirect to login page
	c.SetCookie(users.CookieName, "", -1, "/", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/")

	
}
