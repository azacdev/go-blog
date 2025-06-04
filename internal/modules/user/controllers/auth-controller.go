package controllers

import (
	"log"
	"net/http"

	"github.com/azacdev/go-blog/internal/modules/user/request/auth"
	userService "github.com/azacdev/go-blog/internal/modules/user/services"
	"github.com/azacdev/go-blog/pkg/errors"
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

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "The user has been created successfully",
		"user":    user,
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

	log.Printf("The user has been logged in successfully with name %s \n", user.Name)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User logged in successfully",
		"user":    user,
	})

}

func (controller *Controller) HandleLogout(c *gin.Context) {
	// Invalidate the refresh token in the database
	userID, exists := c.Get("userID") // Assuming a middleware sets this
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Unauthorized: User ID not found in context.",
		})
		return
	}

	err := controller.userService.RevokeRefreshToken(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to logout: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// HandleRefreshToken handles the request to refresh access tokens
func (controller *Controller) HandleRefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ValidationErrorResponse(c, err)
		return
	}

	newAccessToken, newRefreshToken, err := controller.userService.RefreshTokens(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"message":       "Tokens refreshed successfully",
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
