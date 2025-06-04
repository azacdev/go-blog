package services

import (
	"github.com/azacdev/go-blog/internal/modules/user/request/auth"
	userResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (userResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (userResponse.User, error)
	RefreshTokens(refreshToken string) (string, string, error)
	RevokeRefreshToken(userID uint) error
}
