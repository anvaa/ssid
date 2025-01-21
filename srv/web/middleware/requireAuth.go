package middleware

import (
	"srv/srv_sec"
	"usr"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const userKey string = "user"

func RequireAuth(c *gin.Context) {
	// Get the JWT string from the header
	tokenString, err := c.Cookie(users.CookieName)
	if err != nil {
		OnErr(c)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(srv_sec.JwtSecret), nil
	},
	)

	if err != nil {
		OnErr(c)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		OnErr(c)
		return
	}

	user, exists := users.User_GetById(claims["sub"])
	if exists != nil {
		OnErr(c)
		return
	}

	// Attach the user to the context
	c.Set(userKey, user)
	c.Next()
}
