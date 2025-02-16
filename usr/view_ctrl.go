package users
import (
	"app/app_conf"
	"srv/srv_conf"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func Root(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func Version(c *gin.Context) {
	c.String(http.StatusOK, "%s", app_conf.AppInfo())
}

func Info(c *gin.Context) {
	appinfo := app_conf.AppInfo()
	title := fmt.Sprintf("%s v%s", appinfo.AppName, appinfo.Version)

	c.HTML(http.StatusOK, "info.html", gin.H{
		"title":   title,
		"css":     "index.css",
		"url":     c.Request.Referer(),
		"info":    appinfo,
		"company": appinfo.AppName,
		"loctime": app_conf.GetLocalTime(),
		"apptime": app_conf.RunTime(),
		"backbtn": "/app",
	})
}

func View_Signup(c *gin.Context) {
	if c.Request.Method == "POST" {
		SignUp(c)
		return
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{
		"title": "Signup",
		"css":   "index.css",
		"js":    "index.js",
	})
}

func View_Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		Login(c)
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
		"css":   "index.css",
		"js":    "index.js",
	})
}

func View_UserHome(c *gin.Context) {
	uid := c.Keys["user"].(Users).ID

	user_url, err := User_GetUrlFromId(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user url"})
		return
	}

	act, err := User_GetActFromId(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user act"})
		return
	}

	c.HTML(http.StatusOK, "user_home.html", gin.H{
		"title": "Home",
		"user":  c.Keys["user"],
		"css":   "user.css",
		"js":    "user_home.js",
		"url":   user_url,
		"act":   act,
	})
}

func View_NewUsers(c *gin.Context) {
	newusers, cnewusers, _ := Users_GetNew()

	c.HTML(http.StatusOK, "users_new.html", gin.H{
		"title":    "New users",
		"user":     c.Keys["user"],
		"css":      "user.css",
		"js":       "users.js",
		"ginmode":  srv_conf.GetString("gin_mode"),
		"newusers": newusers,
		"countnew": cnewusers,
	})
}

func View_ManageUsers(c *gin.Context) {
	authusers, cauth, _ := Users_GetAuth()
	unauthusers, cunauth, _ := Users_GetUnAuth()
	delusers, cdel, _ := Users_GetDeleted()

	c.HTML(http.StatusOK, "users.html", gin.H{
		"title":       "Manage All Users",
		"user":        c.Keys["user"],
		"css":         "user.css",
		"js":          "users.js",
		"ginmode":     srv_conf.GetString("gin_mode"),
		"authusers":   authusers,
		"countauth":   cauth,
		"unauthusers": unauthusers,
		"countunauth": cunauth,
		"delusers":    delusers,
		"countdel":    cdel,
	})
}

func View_EditUser(c *gin.Context) {
	uid := c.Param("id")

	user, err := User_GetById(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user"})
		return
	}

	var user_url string
	var user_act int64

	if uid != "1" {
		user_url, err = User_GetUrlFromId(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user url"})
			return
		}

		user_act, err = User_GetActFromId(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user act"})
			return
		}
	}

	c.HTML(http.StatusOK, "user_edit.html", gin.H{
		"title":   "Edit User",
		"user":    c.Keys["user"],
		"edituid": user,
		"css":     "user.css",
		"js":      "edit_user.js",
		"ginmode": srv_conf.GetString("gin_mode"),
		"url":     user_url,
		"act":     user_act,
	})
}
