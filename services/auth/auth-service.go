package auth

import (
	"portfolio-user-service/repository/auth/models"
)

type Service interface {
	Register(req models.RegisterRequest) error
	Login(email, password string) (string, error)
	UpdateUserDetails(userID uint, input models.UpdateUserDetailInput) (models.UserDetail, error)
	VerifyEmail(email, name string) error
	VerifyRegistrationOTP(name, email, otp string) error
	GetAllUser() ([]models.User, error)
}
