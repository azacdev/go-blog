package middlewares

import (
	"net/http"
	"strconv"

	UserRepository "github.com/azacdev/go-blog/internal/modules/user/repositories"
	"github.com/azacdev/go-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userId, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userId)

		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Next()

	}
}
