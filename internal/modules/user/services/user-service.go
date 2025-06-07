package services

import (
	"errors"
	"log"

	userModel "github.com/azacdev/go-blog/internal/modules/user/models"
	UserRepository "github.com/azacdev/go-blog/internal/modules/user/repositories"
	"github.com/azacdev/go-blog/internal/modules/user/request/auth"
	UserResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
	"github.com/azacdev/go-blog/pkg/utils"
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

func (userService *UserService) HandleGoogleUser(userInfo auth.GoogleUserInfo) (UserResponse.User, error) {
	var response UserResponse.User

	existingUser := userService.userRepository.FindByEmail(userInfo.Email)

	if existingUser.ID == 0 {
		log.Printf("User with email %s not found. Creating a new user.", userInfo.Email)
		newUserModel := userModel.User{
			Name:     userInfo.Name,
			Email:    userInfo.Email,
			Picture:  userInfo.Picture,
			Password: "",
		}
		existingUser = userService.userRepository.Create(newUserModel)
		if existingUser.ID == 0 {
			return response, errors.New("failed to create new user from Google profile")
		}
	} else {

		existingUser.Name = userInfo.Name
		existingUser.Picture = userInfo.Picture
		if err := userService.userRepository.Update(existingUser); err != nil {
			log.Printf("Failed to update user %d with Google profile info: %v", existingUser.ID, err)

		}
	}

	accessToken, refreshToken, err := utils.GenerateTokens(existingUser.ID)
	if err != nil {
		return response, errors.New("failed to generate authentication tokens")
	}

	existingUser.RefreshToken = refreshToken
	if err := userService.userRepository.Update(existingUser); err != nil {
		log.Printf("Failed to update user %d with refresh token: %v", existingUser.ID, err)
		return response, errors.New("failed to save refresh token for user")
	}

	userRes := UserResponse.ToUser(existingUser)
	userRes.AccessToken = accessToken
	userRes.RefreshToken = refreshToken

	log.Printf("User %s (%d) logged in successfully via Google.", userRes.Email, userRes.ID)

	return userRes, nil
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		log.Printf("Failed to hash password: %v", err) // Use Printf for error
		return response, errors.New("error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)
	// Temporarily create user without specific refresh token (will be generated with ID)
	user.RefreshToken = ""

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error creating the user")
	}

	// Now that we have the user ID, generate proper tokens
	accessToken, refreshToken, err := utils.GenerateTokens(newUser.ID)
	if err != nil {
		log.Printf("Failed to generate tokens for new user %d: %v", newUser.ID, err)

		return response, errors.New("failed to generate authentication tokens")
	}

	// Update the user with the generated refresh token
	newUser.RefreshToken = refreshToken

	if err := userService.userRepository.Update(newUser); err != nil {
		log.Printf("Failed to update user %d with refresh token: %v", newUser.ID, err)

		return response, errors.New("failed to save refresh token for user")
	}

	userRes := UserResponse.ToUser(newUser)
	userRes.AccessToken = accessToken
	userRes.RefreshToken = refreshToken
	return userRes, nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)

	if user.ID != 0 {
		return true
	}

	return false
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User
	existUser := userService.userRepository.FindByEmail(request.Email)

	if existUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	// Generate new access and refresh tokens on successful login
	accessToken, refreshToken, err := utils.GenerateTokens(existUser.ID)
	if err != nil {
		return response, err
	}

	// Update the user's refresh token in the database
	existUser.RefreshToken = refreshToken
	// existUser.TokenExpiration = time.Now().Add(time.Hour * 24 * 7) // If tracking expiration in DB
	if err := userService.userRepository.Update(existUser); err != nil {
		log.Printf("Failed to update user %d with new refresh token: %v", existUser.ID, err)
		return response, errors.New("failed to save new refresh token")
	}

	// Add tokens to the user response
	userRes := UserResponse.ToUser(existUser)
	userRes.AccessToken = accessToken
	userRes.RefreshToken = refreshToken
	return userRes, nil
}

// New service method to handle token refresh
func (userService *UserService) RefreshTokens(refreshToken string) (string, string, error) {
	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid or expired refresh token")
	}

	user := userService.userRepository.FindByID(int(claims.UserID))
	// Critical: Check if the refresh token in the DB matches the one presented
	if user.ID == 0 || user.RefreshToken == "" || user.RefreshToken != refreshToken {
		return "", "", errors.New("refresh token not found, invalid, or revoked")
	}

	// Generate new access and refresh tokens (token rotation)
	newAccessToken, newRefreshToken, err := utils.GenerateTokens(user.ID)
	if err != nil {
		return "", "", err
	}

	// Invalidate the old refresh token and save the new one in the database
	user.RefreshToken = newRefreshToken
	// user.TokenExpiration = time.Now().Add(time.Hour * 24 * 7) // If tracking expiration in DB
	if err := userService.userRepository.Update(user); err != nil {
		log.Printf("Failed to update user %d with rotated refresh token: %v", user.ID, err)
		return "", "", errors.New("failed to save new refresh token during rotation")
	}

	return newAccessToken, newRefreshToken, nil
}

// RevokeRefreshToken clears the refresh token for a given user ID
func (userService *UserService) RevokeRefreshToken(userID uint) error {
	user := userService.userRepository.FindByID(int(userID))
	if user.ID == 0 {
		return errors.New("user not found")
	}

	user.RefreshToken = "" // Clear the refresh token
	// user.TokenExpiration = time.Time{} // Clear expiration as well
	if err := userService.userRepository.Update(user); err != nil {
		log.Printf("Failed to clear refresh token for user %d: %v", user.ID, err)
		return errors.New("failed to revoke refresh token")
	}
	return nil
}
