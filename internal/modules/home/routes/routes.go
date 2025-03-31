package routes

import (
	"net/http"

	"github.com/azacdev/go-blog/pkg/html"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {

		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home Page",
		})
	})

}
