package middlewares

import (
	"net/http"
	"strings"

	UserRepository "github.com/azacdev/go-blog/internal/modules/user/repositories"
	"github.com/azacdev/go-blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	// Initialize the user repository once per middleware instance
	userRepo := UserRepository.New()

	return func(c *gin.Context) {
		// Get the Authorization header from the request
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Authorization header required",
			})
			c.Abort() // Important: Stop further processing
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Invalid Authorization header format. Expected 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		accessToken := parts[1]

		// Validate the access token using your utils function
		claims, err := utils.ValidateAccessToken(accessToken)
		if err != nil {
			// Token is invalid or expired
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Invalid or expired access token: " + err.Error(),
			})
			c.Abort()
			return
		}

		user := userRepo.FindByID(int(claims.UserID))
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "User associated with token not found or is inactive.",
			})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
