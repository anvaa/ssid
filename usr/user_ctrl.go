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

var (
	CookieName = strings.Replace(app_conf.AppName, " ", "", -1) + "_Auth"
	errmsg     = "user or password invalid"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		Password2 string `json:"password2"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to read body"})
		return
	}

	email, password, password2 := strings.TrimSpace(body.Email), strings.TrimSpace(body.Password), strings.TrimSpace(body.Password2)

	if err := IsValidEmail(email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if password != password2 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	if err := IsValidPassword(password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user, err := User_GetByEmail(email)
	if err != nil {
		log.Println("New User")
	}

	if user.Id > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	role, isauth, accessTime, url := "user", false, 3600, app_conf.GetString("usr_url")
	if Users_Count() == 0 {
		role, isauth, accessTime, url = "admin", true, 10800, "/v/newusers"
	}

	user = Users{
		Id:         NewUuid(),
		Email:      email,
		Password:   string(hashedPassword),
		Role:       role,
		IsAuth:     isauth,
		AccessTime: accessTime,
	}

	if err := CreateNewUser(&user, url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "url": "/login"})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read"})
		return
	}

	email, password := strings.TrimSpace(body.Email), strings.TrimSpace(body.Password)
	url := "/"

	if err := IsValidEmail(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg, "url": url})
		return
	}

	if err := IsValidPassword(password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg, "url": url})
		return
	}

	user, err := User_GetByEmail(email)
	if err != nil || !user.IsAuth || user.Email == "" || !CheckPasswordHash(password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg, "url": url})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Second * time.Duration(user.AccessTime)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(srv_sec.JwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(CookieName, tokenString, user.AccessTime, "/", "", true, true)

	if user.Id == 1 {
		url = "/v/newusers"
	} else {
		url, _ = User_GetUrlFromId(user.Id)
	}

	c.JSON(http.StatusMovedPermanently, gin.H{"url": url})
}

func GetAllUsers(c *gin.Context) {
	users, err := Users_GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := User_GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func User_DeleteUser(c *gin.Context) {
	var body struct {
		Id string `json:"id"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	if body.Id == "1212090603" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Not allowed to delete admin!"})
		return
	}

	if err := User_Delete(body.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func User_UpdateAuth(c *gin.Context) {
	var body struct {
		UserId string `json:"id"`
		Isauth string `json:"isauth"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	if body.UserId == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can´t change auth superadmin!"})
		return
	}

	isAuth := body.Isauth == "false"
	if err := user_UpdateAuth(body.UserId, isAuth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update auth"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func User_UpdateRole(c *gin.Context) {
	var body struct {
		Id   string `json:"id"`
		Role string `json:"role"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	if body.Id == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can´t change role superadmin!"})
		return
	}

	if err := user_UpdateRole(body.Id, body.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func User_SetNewPassword(c *gin.Context) {
	var body struct {
		Id       string `json:"id"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	if err := IsValidPassword(body.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Password is not valid"})
		return
	}

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	if err := user_SetNewPassword(body.Id, string(hashedPassword)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func User_SetAct(c *gin.Context) {
	var body struct {
		Id         string `json:"id"`
		AccessTime string `json:"accesstime"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	accessTime := global.CalculateAccessTime(body.AccessTime)
	if err := user_SetAct(body.Id, accessTime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update access time"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func User_UpdateUrl(c *gin.Context) {
	var body struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to read body"})
		return
	}

	if body.Id == "1" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can´t change URL superadmin!"})
		return
	}

	if err := user_UpdateUrl(body.Id, body.Url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
