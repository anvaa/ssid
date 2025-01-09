package app_ctrl

import (
	"app/app_db"
	"app/app_models"
	
	"srv/global"

	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func Loc_AddUpdate(c *gin.Context) {

	var body struct {
		Id  string `json:"id"`
		Txt string `json:"txt"`
		Url string `json:"url"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	body.Txt = strings.TrimSpace(body.Txt)
	id := global.StringToInt(body.Id)

	err := loc_AddUpd(app_models.LocNames{Id: id, Locname: body.Txt})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url})

}

func Loc_Delete(c *gin.Context) {
	var body struct {
		Id  string `json:"id" binding:"required"`
		Url string `json:"url"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := loc_Delete(body.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url})

}

func Loc_GetLocations() ([]app_models.LocNames, int) {
	var locs []app_models.LocNames
	app_db.AppDB.Order("locname asc").Find(&locs)

	return locs, len(locs)
}

func Loc_GetLocName(itmid any) string {
	return app_db.Loc_GetLocName(itmid)
}

func loc_AddUpd(loc app_models.LocNames) error {
	// create or update
	if loc.Id == 0 {
		loc.Id = app_db.Itm_NewItmId()
		err := app_db.AppDB.Create(&loc).Error
		if err != nil {
			return err
		}
	} else {
		err := app_db.AppDB.Model(&loc).Where("id = ?", loc.Id).Update("locname", loc.Locname).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func loc_Delete(id any) error {
	// search items for loc. If found, return error
	var itm app_models.Items
	err := app_db.AppDB.Where("locid = ?", id).First(&itm).Error
	if err == nil {
		newerr := errors.New("location is in use")
		return newerr
	}

	// delete loc
	err = app_db.AppDB.Where("id = ?", id).Delete(&app_models.LocNames{}).Error
	return err
}
