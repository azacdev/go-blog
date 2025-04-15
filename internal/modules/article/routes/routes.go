package routes

import (
	articlesCtrl "github.com/azacdev/go-blog/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := articlesCtrl.New()
	router.GET("/articles/:id", articlesController.Show)
}
