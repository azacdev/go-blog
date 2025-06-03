package routes

import (
	"github.com/azacdev/go-blog/internal/middlewares"
	articlesCtrl "github.com/azacdev/go-blog/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := articlesCtrl.New()
	router.GET("/articles/:id", articlesController.Show)

	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/store", articlesController.Store)
	}
}
