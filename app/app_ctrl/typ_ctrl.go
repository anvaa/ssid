package app_ctrl

import (
	"net/http"
	"strings"
	"errors"

	"app/app_db"
	"app/app_models"
	"srv/global"

	"github.com/gin-gonic/gin"
)

func Typ_AddUpd(c *gin.Context) {
	var body struct {
		Id  string `json:"id"`
		Txt string `json:"txt"`
		Url string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	body.Txt = strings.TrimSpace(body.Txt)
	id := global.StringToInt(body.Id)

	if err := typ_AddUpd(app_models.TypNames{Id: id, Typname: body.Txt}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to add or update type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "url": body.Url})
}

func Typ_Delete(c *gin.Context) {
	var body struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	if err := typ_Delete(body.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "url": body.Url})
}

func Typ_GetTypes() ([]app_models.TypNames, int) {
	var types []app_models.TypNames
	app_db.AppDB.Order("typname asc").Find(&types)
	return types, len(types)
}

func Typ_GetTypName(itmid any) string {
	return app_db.Typ_GetTypName(itmid)
}

func typ_AddUpd(typ app_models.TypNames) error {
	if typ.Id == 0 {
		typ.Id = app_db.Itm_NewItmId()
		if err := app_db.AppDB.Create(&typ).Error; err != nil {
			return err
		}
	} else {
		if err := app_db.AppDB.Model(&typ).Where("id = ?", typ.Id).Update("typname", typ.Typname).Error; err != nil {
			return err
		}
	}
	return nil
}

func typ_Delete(id any) error {
	var itm app_models.Items
	if err := app_db.AppDB.Where("typid = ?", id).First(&itm).Error; err == nil {
		return errors.New("item type is in use")
	}

	return app_db.AppDB.Where("id = ?", id).Delete(&app_models.TypNames{}).Error
}
