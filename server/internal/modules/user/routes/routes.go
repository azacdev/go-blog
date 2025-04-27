package routes

import (
	"github.com/azacdev/go-blog/internal/middlewares"
	userCtrl "github.com/azacdev/go-blog/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		// guestGroup.GET("/register", articlesController.Register)
		guestGroup.POST("/register", articlesController.HandleRegister)

		guestGroup.GET("/login", articlesController.Login)
		guestGroup.POST("/login", articlesController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", articlesController.HandleLogout)
	}
}
