package users

import (
	"app/app_conf"
	"srv/global"
	"srv/srv_sec"

	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var CookieName = app_conf.AppName + "_Auth"
var errmsg = "user or password invalid"

func SignUp(c *gin.Context) {

	var body struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		Password2 string `json:"password2"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to read body"})
		return
	}

	// validate the email/password
	email := strings.TrimSpace(body.Email)
	password := strings.TrimSpace(body.Password)
	password2 := strings.TrimSpace(body.Password2)

	err := IsValidEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error()})
		return
	}

	if password != password2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "InternalServerError"})
		return
	}

	err = IsValidPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error()})
		return
	}

	// check if the email is already in use
	user, err := User_GetByEmail(email)
	if err != nil {
		log.Println("New User")
	}

	if user.Id > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "InternalServerError"})
		return
	}

	// hash the password before save
	hashedPassword, err := HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "InternalServerError"})
		return
	}

	// first user gets to be admin
	role := "user"
	isauth := false
	accessTime := 3600 // 1 hour
	url := app_conf.GetString("usr_url")

	if Users_Count() == 0 {
		role = "admin"
		isauth = true
		accessTime = 10800 // 3 hours
		url = "/v/newusers"
	}

	// create a user
	user = Users{
		Id:         NewUuid(),
		Email:      email,
		Password:   string(hashedPassword),
		Role:       role,
		IsAuth:     isauth,
		AccessTime: accessTime,
	}

	err = CreateNewUser(&user, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error()})
		return
	}

	// return
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     "/login"})

}

func Login(c *gin.Context) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read"})
		return
	}

	url := "/" // default return url

	// validate the email/password
	email := strings.TrimSpace(body.Email)
	password := strings.TrimSpace(body.Password)

	var err error
	err = IsValidEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"url":     url})
		return
	}

	err = IsValidPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"url":     url})
		return
	}

	// find the user by email
	_user, err := User_GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errmsg,
			"url":     url})
		return
	}

	if !_user.IsAuth {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"url":     url})
		return
	}

	if _user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"url":     url})
		return
	}

	if !CheckPasswordHash(password, _user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"url":     url})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": _user.Id,
		"exp": time.Now().Add(time.Second * time.Duration(_user.AccessTime)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(srv_sec.JwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(CookieName, tokenString, _user.AccessTime, "/", "", true, true)

	// redirect to userurl or newusers
	if _user.Id == 1 {
		url = "/v/newusers" // admin start page
	} else {
		url, _ = User_GetUrlFromId(_user.Id) // user start page
	}

	// permanently redirect to url
	c.JSON(301, gin.H{
		"url": url})

}

func GetAllUsers(c *gin.Context) {
	_users, err := Users_GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get users"})
		return
	}

	c.JSON(http.StatusOK, _users)
}

func GetUser(c *gin.Context) {

	id := c.Param("id")
	_user, err := User_GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user"})
		return
	}

	c.JSON(http.StatusOK, _user)
}

func User_DeleteUser(c *gin.Context) {
	var body struct {
		Id string `json:"id"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	uid := body.Id

	if uid == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Not allowed to delete superadmin!"})
		return
	}

	err := User_Delete(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func User_UpdateAuth(c *gin.Context) {
	var body struct {
		UserId string `json:"id"`
		Isauth string `json:"isauth"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	// protect superadmin
	if body.UserId == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can´t change auth superadmin!"})
		return
	}

	var isAuth bool
	switch body.Isauth {
	case "true":
		isAuth = false
	case "false":
		isAuth = true
	}
	// fmt.Println("secind", body)
	err := user_UpdateAuth(body.UserId, isAuth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update auth"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})

}

func User_UpdateRole(c *gin.Context) {
	var body struct {
		Id   string `json:"id"`
		Role string `json:"role"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	// protect superadmin
	if body.Id == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can´t change role superadmin!"})
		return
	}

	err := user_UpdateRole(body.Id, body.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})

}

func User_SetNewPassword(c *gin.Context) {
	var body struct {
		Id       string `json:"id"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	if IsValidPassword(body.Password) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Password is not valid"})
		return
	}

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password"})
		return
	}

	err = user_SetNewPassword(body.Id, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})

}

func User_SetAct(c *gin.Context) {
	var body struct {
		Id         string `json:"id"`
		AccessTime string `json:"accesstime"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	accessTime := global.CalculateAccessTime(body.AccessTime)

	err := user_SetAct(body.Id, accessTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update access time"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})

}

func User_UpdateUrl(c *gin.Context) {
	var body struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to read body"})
		return
	}

	// protect superadmin
	if body.Id == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can´t change URL superadmin!"})
		return
	}

	err := user_UpdateUrl(body.Id, body.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update url"})
		return
	}

	err = user_UpdateUrl(body.Id, body.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})

}
