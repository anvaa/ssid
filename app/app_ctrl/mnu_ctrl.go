package app_ctrl

import (
	
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"app/app_menu"
	"srv/global"
)

func Mnu_UpdTitels(c *gin.Context) {
	var body struct {
		Idx   string `json:"idx" binding:"required"`
		Title string `json:"title" binding:"required"`
		Url   string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	body.Title = strings.TrimSpace(body.Title)
	idx := global.StringToInt(body.Idx)

	app_menu.UpdMenuTitle(idx, body.Title)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     body.Url,
	})

}
