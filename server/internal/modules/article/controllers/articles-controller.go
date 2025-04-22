package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/azacdev/go-blog/internal/modules/article/request/articles"
	ArticleService "github.com/azacdev/go-blog/internal/modules/services"
	"github.com/azacdev/go-blog/internal/modules/user/helpers"
	"github.com/azacdev/go-blog/pkg/converters"
	"github.com/azacdev/go-blog/pkg/errors"
	"github.com/azacdev/go-blog/pkg/html"
	"github.com/azacdev/go-blog/pkg/old"
	"github.com/azacdev/go-blog/pkg/sessions"
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
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// If the article is found, render artilce template
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article", "article": article})

}

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{"title": "Create article"})

}
func (controller *Controller) Store(c *gin.Context) {
	// Validate the request
	var request articles.StoreRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromError(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.URLValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	// Create the article
	article, err := controller.articleService.StoreAsUser(request, user)

	// Check if there is an error on article creation
	if err != nil {
		c.JSON(http.StatusFound, "/article/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
