package app_api

import (
	"log"
	"srv/srv_conf"
	"srv/web/middleware"

	"app/app_ctrl"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func App_Api(r *gin.Engine) *gin.Engine {

	// SET app paths
	static_dir := srv_conf.StaticDir
	log.Println("Static App folder:", static_dir)

	r.Static("/css", static_dir+"/css")
	r.Static("/js", static_dir+"/js")
	r.Static("/media", static_dir+"/media")

	r.Static("/assets", srv_conf.AssetsDir)
	r.Static("/qr_img", srv_conf.QRImgDir)
	r.Static("/bar_img", srv_conf.BarImgDir)

	r.LoadHTMLGlob(static_dir+"/html/*.html")

	// SET app routes
	invGrp := r.Group("/app")
	{
		invGrp.Use(middleware.RequireAuth)
		invGrp.Use(middleware.IsAuth)

		invGrp.GET("/", app_ctrl.Inv_Start)
		invGrp.GET("/home", app_ctrl.Inv_Home)
		invGrp.GET("/search", app_ctrl.Inv_Search)
		invGrp.GET("/tools", app_ctrl.Inv_Tools)
		invGrp.GET("/stats", app_ctrl.Inv_Stats)

	}

	locGrp := r.Group("/loc")
	{
		locGrp.Use(middleware.RequireAuth)
		locGrp.Use(middleware.IsAuth)

		locGrp.POST("/addupd", app_ctrl.Loc_AddUpdate)
		locGrp.DELETE("/delete", app_ctrl.Loc_Delete)

	}

	typGrp := r.Group("/typ")
	{
		typGrp.Use(middleware.RequireAuth)
		typGrp.Use(middleware.IsAuth)

		typGrp.POST("/addupd", app_ctrl.Typ_AddUpd)
		typGrp.DELETE("/delete", app_ctrl.Typ_Delete)

	}

	manGrp := r.Group("/man")
	{
		manGrp.Use(middleware.RequireAuth)
		manGrp.Use(middleware.IsAuth)

		manGrp.POST("/addupd", app_ctrl.Man_AddUpd)
		manGrp.DELETE("/delete", app_ctrl.Man_Delete)

	}

	staGrp := r.Group("/sta")
	{
		staGrp.Use(middleware.RequireAuth)
		staGrp.Use(middleware.IsAuth)

		staGrp.POST("/addupd", app_ctrl.Sta_AddUpd)
		staGrp.DELETE("/delete/:staid", app_ctrl.Sta_HistDelete)

		staGrp.POST("/hist/add", app_ctrl.Sta_HistAdd)

	}

	mnuGrp := r.Group("/menus")
	{
		mnuGrp.Use(middleware.RequireAuth)
		mnuGrp.Use(middleware.IsAuth)

		mnuGrp.POST("/updtitles", app_ctrl.Mnu_UpdTitels)

	}

	itmGrp := r.Group("/itm")
	{
		itmGrp.Use(middleware.RequireAuth)
		itmGrp.Use(middleware.IsAuth)

		itmGrp.GET("/new", app_ctrl.Itm_New)
		itmGrp.POST("/addupd", app_ctrl.Itm_AddUpd)
		itmGrp.DELETE("/delete/:itmid", app_ctrl.Itm_Delete)
		itmGrp.GET("/genqr/:itmid", app_ctrl.Itm_GenQR)

	}

	searchGrp := r.Group("/search")
	{
		searchGrp.Use(middleware.RequireAuth)
		searchGrp.Use(middleware.IsAuth)

		searchGrp.GET("/:itmid", app_ctrl.Itm_GetById)
		searchGrp.GET("/serial/:snr", app_ctrl.Itm_GetBySerial)

		searchGrp.POST("/multi", app_ctrl.Inv_DoSearch)
		searchGrp.POST("/export", app_ctrl.Inv_Export)
	}

	return r

}
