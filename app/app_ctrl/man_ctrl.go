package app_ctrl

import (
	"errors"
	"strings"

	"app/app_db"
	"app/app_models"
	"srv/global"

	"github.com/gin-gonic/gin"
)

func Man_AddUpd(c *gin.Context) {
	var body struct {
		Id  string `json:"id"`
		Txt string `json:"txt" binding:"required"`
		Url string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	body.Txt = strings.TrimSpace(body.Txt)
	id := global.StringToInt(body.Id)

	man := app_models.ManNames{Id: id, Manname: body.Txt}
	if err := manAddOrUpdate(man); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url,
	})
}

func Man_Delete(c *gin.Context) {
	var body struct {
		Id  string `json:"id" binding:"required"`
		Url string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := manDelete(body.Id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url,
	})
}

// DB FUNCTIONS

func Man_GetManufacts() ([]app_models.ManNames, int) {
	var mans []app_models.ManNames
	app_db.AppDB.Order("manname asc").Find(&mans)
	return mans, len(mans)
}

func Man_GetManName(id any) string {
	var manName app_models.ManNames
	app_db.AppDB.Where("id = ?", id).First(&manName)
	return manName.Manname
}

func manAddOrUpdate(man app_models.ManNames) error {
	if man.Id == 0 {
		man.Id = app_db.Itm_NewItmId()
		if err := app_db.AppDB.Create(&man).Error; err != nil {
			return err
		}
	} else {
		if err := app_db.AppDB.Model(&man).Where("id = ?", man.Id).Update("manname", man.Manname).Error; err != nil {
			return err
		}
	}
	return nil
}

func manDelete(id any) error {
	var itm app_models.Items
	if err := app_db.AppDB.Where("manid = ?", id).First(&itm).Error; err == nil {
		return errors.New("manufact is in use")
	}

	return app_db.AppDB.Where("id = ?", id).Delete(&app_models.ManNames{}).Error
}
