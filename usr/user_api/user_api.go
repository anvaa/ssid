package user_api

import (

	"log"

	"github.com/gin-gonic/gin"

	"srv/web/middleware"
	"srv/srv_conf"

	"usr"

)

func User_Api(r *gin.Engine) *gin.Engine {

	// SET app paths
	static_dir := srv_conf.StaticDir
	log.Println("Static App folder:", static_dir)
	r.Static("/media", static_dir+"/media")

	// SET user paths
	r.GET("/", users.Root)
	r.GET("/info", users.Info)
	r.GET("/version", users.Version)

	r.POST("/signup", users.View_Signup)
	r.GET("/signup", users.View_Signup)
	r.POST("/login", users.View_Login)
	r.GET("/login", users.View_Login)
	
	r.GET("/logout", middleware.Logout)

	userRoutes := r.Group("/user")
	{
		userRoutes.Use(middleware.RequireAuth)
		userRoutes.Use(middleware.IsAdmin)
		userRoutes.Use(middleware.IsAuth)

		userRoutes.GET("/", users.GetAllUsers)
		userRoutes.GET("/:id", users.GetUser)

		userRoutes.POST("/delete/:id", users.User_DeleteUser)
		userRoutes.POST("/auth", users.User_UpdateAuth)
		userRoutes.POST("/role", users.User_UpdateRole)
		userRoutes.POST("/psw", users.User_SetNewPassword)
		userRoutes.POST("/act", users.User_SetAct)
		userRoutes.POST("/url", users.User_UpdateUrl)
	}

	viewRoutes := r.Group("/v")
	{
		viewRoutes.Use(middleware.RequireAuth)
		viewRoutes.Use(middleware.IsAuth)
		viewRoutes.Use(middleware.IsAdmin)

		// is admin
		viewRoutes.GET("/newusers", users.View_NewUsers)
		viewRoutes.GET("/users", users.View_ManageUsers)
		viewRoutes.GET("/user/:id", users.View_EditUser)

	}

	return r
}
