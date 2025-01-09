package server

import (
	"users/user_api"

	"app/app_api"

	"srv/srv_conf"

	"github.com/gin-gonic/gin"
)

func InitWebServer() *gin.Engine {

	// GET gin mode from app.yaml
	if srv_conf.IsGinModDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Set up the users routes
	r = user_api.User_Api(r) // sets the routes for the users package

	// Set up the app routes
	r = app_api.App_Api(r) // sets the routes for the app package

	return r
}
