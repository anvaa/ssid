package app_ctrl

import (
	"app/app_conf"
	"app/app_db"
	"app/app_models"

	"srv/global"

	"strconv"

	"github.com/gin-gonic/gin"
)

func Itm_GetItems() ([]app_models.ItemsWeb, int) {
	// Get all items
	itms := []app_models.Items{}

	app_db.AppDB.Order("locid, typid").Find(&itms)
	txtlen := app_conf.TxtLength()

	// Create a slice of ItemsWeb
	itmws := []app_models.ItemsWeb{}
	for _, itm := range itms {

		itmw := app_models.ItemsWeb{
			Itmid:       strconv.Itoa(itm.Itmid),
			Description: global.ShortenText(itm.Description, txtlen),
			Price:       strconv.FormatFloat(itm.Price, 'f', 2, 64),
			Serial:      itm.Serial,
			Updtime:     itm.UpdatedAt.Format("2006-01-02"),

			Loc: Loc_GetLocName(itm.Locid),
			Typ: Typ_GetTypName(itm.Typid),
			Man: Man_GetManName(itm.Manid),
			Sta: Sta_GetStatName(itm.Staid),
			Uid: strconv.Itoa(itm.UserId),
		}
		itmws = append(itmws, itmw)
	}

	return itmws, len(itmws)
}

func Itm_GetItemById(itmid any) app_models.Items {
	// get itm where itmid = itmid
	itm := app_models.Items{}
	app_db.AppDB.Where("itmid = ?", itmid).First(&itm)

	return itm
}

func Itm_GetItemIdBySerial(snr string) int {
	// get itm where itmid = itmid
	itm := app_models.Items{}
	err := app_db.AppDB.Where("serial = ?", snr).First(&itm).Error
	if err != nil {
		return 0
	}

	return itm.Itmid
}

func Itm_AddUpd(c *gin.Context) {
	var body app_models.Items

	// Bind the request body to the Items model
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Add or update the item
	err = app_db.Itm_AddUpd(body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"url": "/app"})

}

func Itm_Delete(c *gin.Context) {
	// Get the item ID
	itmid := c.Param("itmid")
	
	if itmid == "" {
		c.JSON(400, gin.H{"error": "missing itmid"})
		return
	}
	
	// Delete the item
	err := app_db.Itm_Delete(itmid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"url": "/app"})
}

func Itm_GenQR(c *gin.Context) {
	// Get the item ID
	itmid, err := strconv.Atoi(c.Param("itmid"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get the item
	itm := Itm_GetItemById(itmid)

	// Generate the QR code
	err = app_db.Itm_MakeQRCode(itm)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"url": "/app"})
}
