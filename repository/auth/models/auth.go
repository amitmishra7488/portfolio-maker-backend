package models

// User model mapped to "users" table
type User struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName      string `gorm:"type:varchar(100);not null" json:"fullName"`
	Email         string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password      string `gorm:"type:text;not null" json:"password"`
	EmailVerified bool   `gorm:"default:false" json:"emailVerified"`
}

// Input DTO for registration
type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

// LoginRequest represents the login credentials
type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"mypassword"`
}
