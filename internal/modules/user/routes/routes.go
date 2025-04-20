package routes

import (
	userCtrl "github.com/azacdev/go-blog/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := userCtrl.New()

	router.GET("/register", articlesController.Register)
	router.POST("/register", articlesController.HandleRegister)

	router.GET("/login", articlesController.Login)
	router.POST("/login", articlesController.HandleLogin)
}
