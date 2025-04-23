package controllers

import (
	"net/http"

	ArticleService "github.com/azacdev/go-blog/internal/modules/services"
	"github.com/azacdev/go-blog/pkg/errors"
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

	featured, err := controller.articleService.GetFeaturedArticles()

	if err != nil {
		errors.ValidationErrorResponse(c, err) // Use the error handler
		return
	}

	stories, err := controller.articleService.GetStoriesArticles()
	if err != nil {
		errors.ValidationErrorResponse(c, err) // Use the error handler
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully retrieved articles",
		"result": gin.H{
			"featured": featured.Data,
			"stories":  stories.Data,
		},
	})

}
