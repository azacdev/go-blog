package controllers

import (
	"net/http"

	ArticleService "github.com/azacdev/go-blog/internal/modules/services"
	"github.com/azacdev/go-blog/pkg/html"
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

	featured := controller.articleService.GetFeaturedArticles()
	stories := controller.articleService.GetStoriesArticles()

	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page",
		"featured": featured.Data,
		"stories":  stories.Data,
	})

}
