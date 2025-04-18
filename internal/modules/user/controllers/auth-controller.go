package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/azacdev/go-blog/internal/modules/user/request/auth"
	userService "github.com/azacdev/go-blog/internal/modules/user/services"
	"github.com/azacdev/go-blog/pkg/converters"
	"github.com/azacdev/go-blog/pkg/errors"
	"github.com/azacdev/go-blog/pkg/html"
	"github.com/azacdev/go-blog/pkg/old"
	"github.com/azacdev/go-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService userService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: userService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register page",
	})

}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// Validate the request

	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromError(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.URLValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exist")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.URLValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	// Create the user
	user, err := controller.userService.Create(request)

	if err != nil {
		c.JSON(http.StatusFound, "/register")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))
	// Check if there is any error in the user creation
	// After creating the user redirect to homepage
	log.Printf("The user has been created successfully with name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}
