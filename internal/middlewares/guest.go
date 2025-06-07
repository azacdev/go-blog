package middlewares

import (
	"net/http"
	"strings"

	"github.com/azacdev/go-blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {

			c.Next()
			return
		}

		accessToken := parts[1]
		_, err := utils.ValidateAccessToken(accessToken)
		if err == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Already authenticated",
			})
		}

		c.Next()
	}
}
