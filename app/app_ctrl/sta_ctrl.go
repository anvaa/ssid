package app_ctrl

import (
	"fmt"
	"strconv"
	"strings"

	"app/app_db"
	"app/app_models"
	"srv/global"
	users "usr"

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

	if err := sta_AddUpd(app_models.StaNames{Id: id, Staname: body.Txt}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success", "url": body.Url})
}

func Sta_HistDelete(c *gin.Context) {

	staid := c.Param("staid")

	if staid == "" {
		c.JSON(400, gin.H{"error": "missing id"})
		return
	}

	if err := staHist_Delete(staid); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})

}

func sta_GetStatHistory(itmid any) ([]app_models.ItmStatHistWeb, error) {
	var stath []app_models.Status_History
	if err := app_db.AppDB.Where("itmid = ?", itmid).Order("updated_at desc").Find(&stath).Error; err != nil {
		return nil, err
	}

	var staHistWeb []app_models.ItmStatHistWeb
	for _, s := range stath {
		updTime := s.UpdatedAt.Format("2006-01-02")
		uid_email, err := users.User_GetEmailById(s.UserId)
		if err != nil {
			uid_email = strconv.Itoa(s.UserId)
		}

		staHistWeb = append(staHistWeb, app_models.ItmStatHistWeb{
			Id:      s.Id,
			Staid:   s.Staid,
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
	if sta.Id == 0 {
		sta.Id = app_db.Itm_NewItmId()
		if err := app_db.AppDB.Create(&sta).Error; err != nil {
			return err
		}
	} else {
		if err := app_db.AppDB.Model(&sta).Where("id = ?", sta.Id).Update("staname", sta.Staname).Error; err != nil {
			return err
		}
	}
	return nil
}

func staHist_Delete(id any) error {
	// delete from status_history if exists
	var stah app_models.Status_History
	if err := app_db.AppDB.Where("id = ?", id).First(&stah).Error; err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}

	// delete from status_history
	if err := app_db.AppDB.Delete(&stah).Error; err != nil {
		return err
	}

	return nil

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

	if err := sta_HistAdd(app_models.Status_History{
		Itmid:   body.Itmid,
		Staid:   body.Staid,
		Comment: body.Txt,
		UserId:  body.Uid,
	}); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success", "url": "/search/" + global.IntToString(body.Itmid)})
}

func sta_HistAdd(newhist app_models.Status_History) error {
	stahist := app_models.Status_History{
		Itmid:   newhist.Itmid,
		Staid:   newhist.Staid,
		UserId:  newhist.UserId,
		Comment: newhist.Comment,
	}

	if err := app_db.AppDB.Create(&stahist).Error; err != nil {
		return err
	}

	if err := app_db.Itm_UpdCurStatus(newhist.Itmid, newhist.Staid); err != nil {
		return err
	}

	return nil
}

func Sta_GetStatName(id any) string {
	return app_db.Sta_GetStaName(id)
}
