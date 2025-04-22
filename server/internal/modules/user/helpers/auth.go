package helpers

import (
	"strconv"

	UserRepository "github.com/azacdev/go-blog/internal/modules/user/repositories"
	UserResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
	"github.com/azacdev/go-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	var userRepo = UserRepository.New()

	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return response
	}

	return UserResponse.ToUser(user)
}
