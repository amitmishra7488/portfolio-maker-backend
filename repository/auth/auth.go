package auth

import (
	"errors"
	"fmt"
	"time"

	"portfolio-user-service/config"
	"portfolio-user-service/repository/auth/models"

	"gorm.io/gorm"
)

type AuthRepository struct{}

func (ar AuthRepository) GetUserByEmailTx(tx *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar AuthRepository) CreateUserTx(tx *gorm.DB, user *models.User) error {
	return tx.Create(user).Error
}

func (ar AuthRepository) CreateUserProfileTx(tx *gorm.DB, profile *models.UserDetail) error {
	return tx.Create(profile).Error
}

// GetUserByEmailQuery fetches a user by email
func (ar AuthRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

// User Details
func (ar AuthRepository) GetUserDetailsByUserID(userID uint) (*models.UserDetail, error) {
	var detail models.UserDetail
	if err := config.DB.Where("user_id = ?", userID).First(&detail).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

// update user details
func (ar AuthRepository) UpdateUserDetails(detail *models.UserDetail) error {
	if err := config.DB.Save(detail).Error; err != nil {
		return err
	}
	return nil
}

// VerifyUserEmailAddress verifies the user's email address
func (ar AuthRepository) VerifyUserEmailAddress(email string) error {
	result := config.DB.Model(&models.User{}).
		Where("email = ? AND email_verified = false", email).
		Updates(map[string]interface{}{
			"email_verified":    true,
			"email_verified_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no unverified user found with email: %s", email)
	}

	return nil
}
