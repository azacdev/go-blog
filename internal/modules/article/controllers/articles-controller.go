package controllers

import (
	"net/http"
	"strconv"

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

func (controller *Controller) Show(c *gin.Context) {
	// Get the article id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "Error converting the id"})
		return
	}

	// Find the artlcle from the database
	article, err := controller.articleService.Find(id)

	// If the article is not found show error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// If the article is found, render artilce template
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article", "article": article})
	// c.JSON(http.StatusOK, gin.H{"article": article})
}
