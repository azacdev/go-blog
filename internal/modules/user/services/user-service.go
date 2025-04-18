package services

import (
	"errors"
	"log"

	userModel "github.com/azacdev/go-blog/internal/modules/user/models"
	UserRepository "github.com/azacdev/go-blog/internal/modules/user/repositories"
	"github.com/azacdev/go-blog/internal/modules/user/request/auth"
	UserResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)

	if err != nil {
		log.Fatal("Failed to hash password")
		return response, errors.New("error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error creating the user")
	}

	return UserResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)

	if user.ID != 0 {
		return true
	}

	return false
}
