package middlewares

import (
	"log"
	"net/http"
	"strconv"

	UserRepository "github.com/azacdev/go-blog/internal/modules/user/repositories"
	"github.com/azacdev/go-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userId, _ := strconv.Atoi(authID)

		log.Printf("User ID: %d", userId)

		user := userRepo.FindByID(userId)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Next()

	}
}
