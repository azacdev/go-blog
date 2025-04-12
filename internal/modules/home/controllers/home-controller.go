package controllers

import (
	"net/http"

	// ArticleRepository "github.com/azacdev/go-blog/internal/modules/article/repositories"
	ArticleService "github.com/azacdev/go-blog/internal/modules/services"
	// "github.com/azacdev/go-blog/pkg/html"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {

	// html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	// 	"title": "Home Page",
	// })

	c.JSON(http.StatusOK, gin.H{
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}
