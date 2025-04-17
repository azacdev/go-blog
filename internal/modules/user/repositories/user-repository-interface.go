package repositories

import userModel "github.com/azacdev/go-blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
}
