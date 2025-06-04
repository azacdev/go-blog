package controllers

import (
	"net/http"
	"strconv"

	"github.com/azacdev/go-blog/internal/modules/article/request/articles"
	ArticleService "github.com/azacdev/go-blog/internal/modules/services"
	"github.com/azacdev/go-blog/internal/modules/user/helpers"
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

func (controller *Controller) Show(c *gin.Context) {
	// Get the article id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errors.ValidationErrorResponse(c, err)
		return
	}

	// Find the artlcle from the database
	article, err := controller.articleService.Find(id)

	if err != nil {
		errors.ValidationErrorResponse(c, err) // Use the error handler
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"article": article,
		"message": "Article fetched successfully",
	})

}

func (controller *Controller) Store(c *gin.Context) {
	// Validate the request
	var request articles.StoreRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ValidationErrorResponse(c, err)
		return
	}

	user, err := helpers.Auth(c)

	if err != nil {
		errors.FieldErrorResponse(c, http.StatusUnauthorized, "Authentication failed: "+err.Error())
		return
	}

	// Create the article
	article, err := controller.articleService.StoreAsUser(request, user)

	// Check if there is an error on article creation
	if err != nil {
		errors.ValidationErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Article has been created successfully",
		"article": article,
	})
}
