package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/azacdev/go-blog/internal/modules/user/request/auth"
	userService "github.com/azacdev/go-blog/internal/modules/user/services"
	"github.com/azacdev/go-blog/pkg/errors"
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

func (controller *Controller) HandleRegister(c *gin.Context) {
	// Validate the request
	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.ValidationErrorResponse(c, err)
		return
	}

	// Check if there is any error in the user creation
	if controller.userService.CheckUserExists(request.Email) {
		errors.FieldErrorResponse(c, http.StatusConflict, "Email address already exists")
		return
	}

	// Create the user
	user, err := controller.userService.Create(request)

	if err != nil {
		errors.ValidationErrorResponse(c, err)
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))
	// After creating the user redirect to homepage
	log.Printf("The user has been created successfully with name %s \n", user.Name)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "The user has been created successfully",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	// Validate the request
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.ValidationErrorResponse(c, err)
		return
	}

	user, err := controller.userService.HandleUserLogin(request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))
	// After creating the user redirect to homepage
	log.Printf("The user has been logged in successfully with name %s \n", user.Name)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User loggedin succesfully",
		"user":    user,
	})

}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
