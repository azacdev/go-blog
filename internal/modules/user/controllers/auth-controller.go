package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/azacdev/go-blog/internal/modules/user/request/auth"
	userService "github.com/azacdev/go-blog/internal/modules/user/services"
	"github.com/azacdev/go-blog/pkg/config"
	"github.com/azacdev/go-blog/pkg/errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Controller struct {
	userService userService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: userService.New(),
	}
}

var googleOauth2Config *oauth2.Config
var oauthStateString = "go-backend"

func init() {
	config.Set()

	cfg := config.Get()

	googleOauth2Config = &oauth2.Config{
		ClientID:     cfg.GoogleOAuth.ClientID,
		ClientSecret: cfg.GoogleOAuth.ClientSecret,
		RedirectURL:  "http://localhost:3000/dashboard/callback/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

}

func (controller *Controller) HandleGoogleLogin(c *gin.Context) {
	url := googleOauth2Config.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func (controller *Controller) HandleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		errors.FieldErrorResponse(c, http.StatusBadRequest, "Invalid OAuth state")
		return
	}

	code := c.Query("code")
	token, err := googleOauth2Config.Exchange(c, code)
	if err != nil {
		errors.FieldErrorResponse(c, http.StatusInternalServerError, "Failed to exchange code for token: "+err.Error())
		return
	}

	client := googleOauth2Config.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		errors.FieldErrorResponse(c, http.StatusInternalServerError, "Failed to get user info: "+err.Error())
		return
	}
	defer resp.Body.Close()

	// Use the dedicated struct from the service layer
	var userInfo auth.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		errors.FieldErrorResponse(c, http.StatusInternalServerError, "Failed to decode user info: "+err.Error())
		return
	}

	user, err := controller.userService.HandleGoogleUser(userInfo)
	if err != nil {
		errors.FieldErrorResponse(c, http.StatusInternalServerError, "Failed to process Google login: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User logged in successfully via Google",
		"user":    user,
	})
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
