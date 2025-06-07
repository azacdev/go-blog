package routes

import (
	"github.com/azacdev/go-blog/internal/middlewares"
	userCtrl "github.com/azacdev/go-blog/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.POST("/register", userController.HandleRegister)
		guestGroup.POST("/login", userController.HandleLogin)
		guestGroup.GET("/login/google", userController.HandleGoogleLogin)
		guestGroup.GET("/login/google/callback", userController.HandleGoogleCallback)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", userController.HandleLogout)
	}

	router.POST("/refresh-token", userController.HandleRefreshToken)
}
