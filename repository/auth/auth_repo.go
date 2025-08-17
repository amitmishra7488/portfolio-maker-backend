package auth

import (
	"portfolio-user-service/repository/auth/models"

	"gorm.io/gorm"
)

// smaller interfaces
type UserReader interface {
    GetAllUser() ([]models.User, error)
    GetUserByEmail(email string) (models.User, error)
	GetUserByEmailTx(tx *gorm.DB, email string) (*models.User, error)
}

type UserWriter interface {
    CreateUserTx(tx *gorm.DB, user *models.User) error
    UpdateUserDetails(detail *models.UserDetail) error
}

type VerificationRepo interface {
    VerifyUserEmailAddress(email string) error
}

// one big interface if needed
type AuthRepo interface {
    UserReader
    UserWriter
    VerificationRepo
    GetUserDetailsByUserID(userID uint) (*models.UserDetail, error)
    CreateUserProfileTx(tx *gorm.DB, profile *models.UserDetail) error
}

