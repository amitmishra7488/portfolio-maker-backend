package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"portfolio-user-service/repository/auth"
	"portfolio-user-service/repository/auth/models"
	"portfolio-user-service/utils"

	"go.uber.org/zap"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuthService struct {
	Repo auth.AuthRepo
	Log  *zap.Logger
	DB   *gorm.DB
}

func NewAuthService(repo auth.AuthRepo, db *gorm.DB, logger *zap.Logger) *AuthService {
	return &AuthService{
		Repo: repo,
		Log:  logger,
		DB:   db,
	}
}

func (s *AuthService) Register(registerRequest models.RegisterRequest) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		// Check if user already exists
		_, err := s.Repo.GetUserByEmailTx(tx, registerRequest.Email)
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
		if err := s.Repo.CreateUserTx(tx, user); err != nil {
			return errors.New("failed to create user")
		}

		// Create empty profile
		profile := &models.UserDetail{
			UserID:   user.ID,
			Username: nil,
		}
		if err := s.Repo.CreateUserProfileTx(tx, profile); err != nil {
			return errors.New("failed to create user profile")
		}

		return nil
	})
}

func (s *AuthService) Login(email, password string) (string, error) {
	// Check if user exists
	user, err := s.Repo.GetUserByEmail(email)
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

func (s *AuthService) UpdateUserDetails(userID uint, input models.UpdateUserDetailInput) (models.UserDetail, error) {
	// 1. Get existing details from repository
	detail, err := s.Repo.GetUserDetailsByUserID(userID)
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
	if err := s.Repo.UpdateUserDetails(detail); err != nil {
		return models.UserDetail{}, errors.New("failed to update user details")
	}

	return *detail, nil
}

func (s *AuthService) VerifyEmail(email, name string) error {
	// Generate OTP
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	subject := "Account Registration"
	message := fmt.Sprintf("Your OTP for account verification is: %s", otp)
	err = utils.SendEmail(email, name, message, subject)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}

	// Save OTP to Redis
	// err = utils.SaveOTP(email, otp)
	// if err != nil {
	// 	return fmt.Errorf("failed to save OTP: %w", err)
	// }

	return nil
}

func (s *AuthService) VerifyRegistrationOTP(name, email, otp string) error {
	// Check if OTP is valid
	isValid, err := utils.CheckOTP(email, otp)
	if err != nil {
		return fmt.Errorf("failed to check OTP: %w", err)
	}
	if !isValid {
		return errors.New("invalid OTP")
	}

	// Update user status to verified
	err = s.Repo.VerifyUserEmailAddress(email)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	subject := "Account Registration SuccessFully"
	message := fmt.Sprintf("Welcome %s! Your email has been successfully verified. You can now access all the features of your account.", name)
	err = utils.SendEmail(email, name, message, subject)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}

	// Delete OTP from Redis
	// err = utils.DeleteOTP(email)
	// if err != nil {
	// 	return fmt.Errorf("failed to delete OTP: %w", err)
	// }

	return nil
}

func (s *AuthService) GetAllUser() ([]models.User, error) {
	var users []models.User
	users, err := s.Repo.GetAllUser()
	if err != nil {
		return users, errors.New("something went wrong" + err.Error())
	}

	return users, nil
}
