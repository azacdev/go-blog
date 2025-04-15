package controllers

import (
	"net/http"
	"strconv"

	ArticleService "github.com/azacdev/go-blog/internal/modules/services"
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Error converting the id"})
		return
	}

	// Find the artlcle from the database

	// If the article is not found show error

	// If the article is found, render artilce template
	c.JSON(http.StatusOK, gin.H{"message": id})
}
