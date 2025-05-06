package auth

import (
	"encoding/json"
	"errors"
	"portfolio-user-service/config"
	"portfolio-user-service/repository/auth"
	"portfolio-user-service/repository/auth/models"
	"portfolio-user-service/utils"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// AuthService struct
type AuthService struct{}

var (
	authRepo = new(auth.AuthRepository)
)

func (s AuthService) Register(registerRequest models.RegisterRequest) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		// Check if user already exists
		_, err := authRepo.GetUserByEmailTx(tx, registerRequest.Email)
		if err == nil {
			return errors.New("email already registered")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err // real DB error
		}

		// Generate full name
		fullName, err := utils.CreateFullName(registerRequest.FirstName, registerRequest.LastName)
		if err != nil {
			return err
		}

		// Hash password
		hashedPassword, err := utils.HashPassword(registerRequest.Password)
		if err != nil {
			return errors.New("failed to hash password")
		}

		// Create user
		user := &models.User{
			FullName: fullName,
			Email:    registerRequest.Email,
			Password: hashedPassword,
		}
		if err := authRepo.CreateUserTx(tx, user); err != nil {
			return errors.New("failed to create user")
		}

		// Create empty profile
		profile := &models.UserDetail{
			UserID: user.ID,
			Username: nil,
		}
		if err := authRepo.CreateUserProfileTx(tx, profile); err != nil {
			return errors.New("failed to create user profile")
		}

		return nil
	})
}


func (s AuthService) Login(email, password string) (string, error) {
	// Check if user exists
	user, err := authRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Check if password matches
	passwordCheck := utils.CheckPasswordHash(password, user.Password)
	if !passwordCheck {
		return "", errors.New("invalid email or password")
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", errors.New("invalid request")
	}

	return token, nil
}

func (s AuthService) UpdateUserDetails(userID uint, input models.UpdateUserDetailInput) (models.UserDetail, error) {
	// 1. Get existing details from repository
	detail, err := authRepo.GetUserDetailsByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.UserDetail{}, errors.New("user details not found")
		}
		return models.UserDetail{}, err
	}

	// 2. Apply only non-nil fields
	if input.Username != nil {
		detail.Username = input.Username
	}
	if input.Bio != nil {
		detail.Bio = *input.Bio
	}
	if input.ProfilePicture != nil {
		detail.ProfilePicture = *input.ProfilePicture
	}
	if input.SocialLinks != nil {
		socialJSON, err := json.Marshal(input.SocialLinks)
		if err != nil {
			return models.UserDetail{}, errors.New("invalid social links format")
		}
		detail.SocialLinks = datatypes.JSON(socialJSON)
	}

	// 3. Save via repository
	if err := authRepo.UpdateUserDetails(detail); err != nil {
		return models.UserDetail{}, errors.New("failed to update user details")
	}

	return *detail, nil
}


