package responses

import (
	"fmt"

	userModels "github.com/azacdev/go-blog/internal/modules/user/models"
)

type User struct {
	ID           uint
	Image        string
	Name         string
	Email        string
	AccessToken  string
	RefreshToken string
}

type Users struct {
	Data []User
}

func ToUser(user userModels.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
