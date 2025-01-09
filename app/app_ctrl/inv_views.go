package app_ctrl

import (
	"app/app_conf"
	"app/app_models"
	"srv/filefunc"
	"srv/srv_conf"

	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var appinfo = app_conf.AppInfo()

func Inv_Start(c *gin.Context) {

	c.HTML(http.StatusOK, "inv_start.html", gin.H{
		"title": app_conf.AppName+" - Home",

		"css": "inv_lists.css",
		"js":  "inv_start.js",

		"user":    c.Keys["user"],
		"appinfo": appinfo,
	})
}

func Inv_Home(c *gin.Context) {

	itmsw, c_itms := Itm_GetItems()

	c.HTML(http.StatusOK, "inv_home.html", gin.H{

		"itmsw":   itmsw,
		"c_itms":  c_itms,
		"user":    c.Keys["user"],
		"appinfo": appinfo,
	})
}

func Inv_Tools(c *gin.Context) {

	locs, c_locs := Loc_GetLocations()
	typs, c_typs := Typ_GetTypes()
	mans, c_mans := Man_GetManufacts()
	stats, c_stats := Stat_GetStatuses()

	c.HTML(http.StatusOK, "inv_tools.html", gin.H{
		"title": app_conf.AppName+" - Tools",

		"css":     "inv_lists.css",
		"js":      "inv_tools.js",
		"user":    c.Keys["user"],
		"appinfo": appinfo,

		"locs":    locs,
		"c_locs":  c_locs,
		"typs":    typs,
		"c_typs":  c_typs,
		"mans":    mans,
		"c_mans":  c_mans,
		"stats":   stats,
		"c_stats": c_stats,
	})
}

func Inv_Stats(c *gin.Context) {

	stats, err := GetStats()
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "inv_stats.html", gin.H{
		"title": app_conf.AppName+" - Statistics",

		"css":     "inv_lists.css",
		"js":      "inv_tools.js",
		"user":    c.Keys["user"],
		"appinfo": appinfo,

		"stats": stats,
	})
}

func Itm_New(c *gin.Context) {

	locs, _ := Loc_GetLocations()
	typs, _ := Typ_GetTypes()
	mans, _ := Man_GetManufacts()

	c.HTML(http.StatusOK, "itm_new.html", gin.H{

		"js":      "itm_edit.js",
		"css":     "inv_lists.css",
		"user":    c.Keys["user"],
		"appinfo": appinfo,

		"locs": locs,
		"typs": typs,
		"mans": mans,
	})
}

func Itm_GetById(c *gin.Context) {

	itmid := c.Param("itmid")

	locs, _ := Loc_GetLocations()
	typs, _ := Typ_GetTypes()
	mans, _ := Man_GetManufacts()
	stat, _ := Stat_GetStatuses()

	itm := Itm_GetItemById(itmid)
	sthist, _ := sta_GetStatHistory(itmid)

	itmw := app_models.ItemsWeb{
		Itmid: strconv.Itoa(itm.Itmid),

		Description: strings.Trim(itm.Description, " "),
		Price:       strconv.FormatFloat(itm.Price, 'f', 2, 64),
		Serial:      itm.Serial,
		Updtime:     itm.UpdatedAt.Format("2006-01-02 15:04:05"),

		Loc: Loc_GetLocName(itm.Locid),
		Typ: Typ_GetTypName(itm.Typid),
		Man: Man_GetManName(itm.Manid),

		Uid: strconv.Itoa(itm.UserId),
	}

	c.HTML(http.StatusOK, "itm_edit.html", gin.H{

		"js":      "itm_edit.js",
		"css":     "inv_lists.css",
		"user":    c.Keys["user"],
		"appinfo": appinfo,

		"locs":  locs,
		"typs":  typs,
		"mans":  mans,
		"stats": stat,
		"itmw":  itmw,
		"statw": sthist,
	})
}

func Itm_GetBySerial(c *gin.Context) {

	snr := c.Param("snr")
	itmid := Itm_GetItemIdBySerial(snr)
	c.Redirect(http.StatusFound, "/search/"+strconv.Itoa(itmid))

}

func Inv_Search(c *gin.Context) {

	locs, _ := Loc_GetLocations()
	typs, _ := Typ_GetTypes()
	mans, _ := Man_GetManufacts()
	stat, _ := Stat_GetStatuses()

	c.HTML(http.StatusOK, "inv_search.html", gin.H{
		"title": app_conf.AppName+" - Search",

		"css":     "inv_lists.css",
		"js":      "inv_search.js",
		"user":    c.Keys["user"],
		"appinfo": appinfo,

		"locs":  locs,
		"typs":  typs,
		"mans":  mans,
		"stats": stat,
	})
}

func Inv_DoSearch(c *gin.Context) {

	var search app_models.DoSearch
	c.BindJSON(&search)

	itmsw, c_itms := inv_DoSearch(search, false)

	c.HTML(http.StatusOK, "itm_search_result.html", gin.H{

		"itmsw":  itmsw,
		"c_itms": c_itms,
	})
}

func Inv_Export(c *gin.Context) {

	var search app_models.DoSearch
	c.BindJSON(&search)

	itmsw, _ := inv_DoSearch(search, true)

	filename := fmt.Sprintf("%s_%s.csv", time.Now().Format("0601021504") , app_conf.AppName)
	filepath := fmt.Sprintf("%s/%s", srv_conf.ReportsDir, filename)

	err := filefunc.ExportSearchResult(filepath, filename, itmsw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// download file
	c.FileAttachment(filepath, filename)

	// Delete the file after the user has downloaded it
	go filefunc.DeleteFile(filepath)

}
