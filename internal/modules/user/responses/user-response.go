package responses

import (
	userModels "github.com/azacdev/go-blog/internal/modules/user/models"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	Picture      string
	AccessToken  string
	RefreshToken string
}

type Users struct {
	Data []User
}

func ToUser(user userModels.User) User {
	return User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
	}
}
