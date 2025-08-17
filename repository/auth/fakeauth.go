package auth

import (
	"portfolio-user-service/repository/auth/models"

	"gorm.io/gorm"
)

// FakeAuthRepo is an in-memory implementation of AuthRepo for testing
type FakeAuthRepo struct {
	Users    map[string]*models.User
	Profiles map[uint]*models.UserDetail
}

// NewFakeAuthRepo creates a new fake repo with empty maps
func NewFakeAuthRepo() *FakeAuthRepo {
	return &FakeAuthRepo{
		Users:    make(map[string]*models.User),
		Profiles: make(map[uint]*models.UserDetail),
	}
}

// GetAllUser returns all users
func (f *FakeAuthRepo) GetAllUser() ([]models.User, error) {
	var users []models.User
	for _, u := range f.Users {
		users = append(users, *u)
	}
	return users, nil
}

// GetUserByEmailTx returns a user by email
func (f *FakeAuthRepo) GetUserByEmailTx(tx *gorm.DB, email string) (*models.User, error) {
	if user, ok := f.Users[email]; ok {
		return user, nil
	}
	return nil, gorm.ErrRecordNotFound
}

// GetUserByEmail returns a user by email
func (f *FakeAuthRepo) GetUserByEmail(email string) (models.User, error) {
	if user, ok := f.Users[email]; ok {
		return *user, nil
	}
	return models.User{}, gorm.ErrRecordNotFound
}

// CreateUserTx adds a user
func (f *FakeAuthRepo) CreateUserTx(tx *gorm.DB, user *models.User) error {
	f.Users[user.Email] = user
	return nil
}

// UpdateUserDetails updates user detail
func (f *FakeAuthRepo) UpdateUserDetails(detail *models.UserDetail) error {
	f.Profiles[detail.UserID] = detail
	return nil
}

// GetUserDetailsByUserID gets user detail by user ID
func (f *FakeAuthRepo) GetUserDetailsByUserID(userID uint) (*models.UserDetail, error) {
	if profile, ok := f.Profiles[userID]; ok {
		return profile, nil
	}
	return nil, gorm.ErrRecordNotFound
}

// CreateUserProfileTx adds a user profile
func (f *FakeAuthRepo) CreateUserProfileTx(tx *gorm.DB, profile *models.UserDetail) error {
	f.Profiles[profile.UserID] = profile
	return nil
}

// VerifyUserEmailAddress checks if a user exists
func (f *FakeAuthRepo) VerifyUserEmailAddress(email string) error {
	if _, ok := f.Users[email]; ok {
		return nil
	}
	return gorm.ErrRecordNotFound
}
