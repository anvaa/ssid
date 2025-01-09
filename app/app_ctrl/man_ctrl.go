package app_ctrl

import (

	"app/app_db"
	"app/app_models"

	"srv/global"

	"strings"
	"errors"

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
	if err := man_AddUpd(man); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url})

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

	if err := man_Delete(body.Id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url})
}

// DB FUNCTIONS

func Man_GetManufacts() ([]app_models.ManNames, int) {
	var mans []app_models.ManNames
	app_db.AppDB.Order("manname asc").Find(&mans)

	return mans, len(mans)
}

func Man_GetManName(id any) string {
	var man_name app_models.ManNames
	app_db.AppDB.Where("id = ?", id).First(&man_name)
	return man_name.Manname
}

func man_AddUpd(man app_models.ManNames) error {
	// create or update
	if man.Id == 0 {
		man.Id = app_db.Itm_NewItmId()
		err := app_db.AppDB.Create(&man).Error
		if err != nil {
			return err
		}
	} else {
		err := app_db.AppDB.Model(&man).Where("id = ?", man.Id).Update("manname", man.Manname).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func man_Delete(id any) error {
	// search items for man. If found, return error
	var itm app_models.Items
	err := app_db.AppDB.Where("manid = ?", id).First(&itm).Error
	if err == nil {
		newerr := errors.New("manufact is in use")
		return newerr
	}

	err = app_db.AppDB.Where("id = ?", id).Delete(&app_models.ManNames{}).Error
	return err
}
