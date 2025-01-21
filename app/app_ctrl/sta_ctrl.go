package app_ctrl

import (
	"srv/global"
	"usr"

	"app/app_db"
	"app/app_models"

	"fmt"
	"strconv"
	"strings"
	"errors"

	"github.com/gin-gonic/gin"
)

func Stat_GetStatuses() ([]app_models.StaNames, int) {
	var stats []app_models.StaNames
	app_db.AppDB.Order("staname asc").Find(&stats)

	return stats, len(stats)
}

func Sta_GetLatestStat(itmid any) string {
	var stah app_models.Status_History
	app_db.AppDB.Where("itmid = ?", itmid).Order("updated_at desc").First(&stah)

	fmt.Println("Sta_GetLatestStat", Sta_GetStatName(stah.Staid))
	return Sta_GetStatName(stah.Staid)

}

func Sta_AddUpd(c *gin.Context) {
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

	err := sta_AddUpd(app_models.StaNames{Id: id, Staname: body.Txt})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url})

}

func Sta_Delete(c *gin.Context) {
	var body struct {
		Id  string `json:"id" binding:"required"`
		Url string `json:"url"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := sta_Delete(body.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     body.Url})

}

func sta_GetStatHistory(itmid any) ([]app_models.ItmStatHistWeb, error) {
	
	var stath []app_models.Status_History
	err := app_db.AppDB.Where("itmid = ?", itmid).
		Order("updated_at desc").
		Find(&stath).Error
	if err != nil {
		return nil, err
	}

	staHistWeb := []app_models.ItmStatHistWeb{}
	for _, s := range stath {

		updTime := s.UpdatedAt.Format("2006-01-02 15:04:05")
		uid_email, err := users.User_GetEmailById(s.UserId)
		if err != nil {
			uid_email = strconv.Itoa(s.UserId)
		}

		staHistWeb = append(staHistWeb, app_models.ItmStatHistWeb{
			Staname: Sta_GetStatName(s.Staid),
			Updated: updTime,
			Uid:     uid_email,
			Comment: s.Comment,
		})
	}

	return staHistWeb, nil
}

func Sta_GetStatNames() ([]app_models.StaNames, int) {
	var stats []app_models.StaNames
	app_db.AppDB.Order("staname").Find(&stats)

	return stats, len(stats)
}

func sta_AddUpd(sta app_models.StaNames) error {
	// create or update
	if sta.Id == 0 {
		sta.Id = app_db.Itm_NewItmId()
		err := app_db.AppDB.Create(&sta).Error
		if err != nil {
			return err
		}
	} else {
		err := app_db.AppDB.Model(&sta).Where("id = ?", sta.Id).Update("staname", sta.Staname).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func sta_Delete(id any) error {
	// search items for sta. If found, return error
	var stah app_models.Status_History
	err := app_db.AppDB.Where("staid = ?", id).First(&stah).Error
	if err == nil {
		newerr := errors.New("status is in use")
		return newerr
	}

	err = app_db.AppDB.Where("id = ?", id).Delete(&app_models.StaNames{}).Error
	return err
}

func Sta_HistAdd(c *gin.Context) {
	var body struct {
		Itmid int    `json:"itmid"`
		Staid int    `json:"staid"`
		Txt   string `json:"txt"`
		Uid   int    `json:"uid"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	body.Txt = strings.TrimSpace(body.Txt)

	err := sta_HistAdd(app_models.Status_History{
		Itmid:   body.Itmid,
		Staid:   body.Staid,
		Comment: body.Txt,
		UserId:  body.Uid,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"url":     "/search/" + global.IntToString(body.Itmid),
	})

}

func sta_HistAdd(newhist app_models.Status_History) error {

	stahist := app_models.Status_History{
		Itmid:   newhist.Itmid,
		Staid:   newhist.Staid,
		UserId:  newhist.UserId,
		Comment: newhist.Comment,
	}

	err := app_db.AppDB.Create(&stahist).Error
	if err != nil {
		return err
	}

	// update the current status of the items table
	err = app_db.Itm_UpdCurStatus(newhist.Itmid, newhist.Staid)
	if err != nil {
		return err
	}

	return nil
}

func Sta_GetStatName(id any) string {
	return app_db.Sta_GetStaName(id)
}
