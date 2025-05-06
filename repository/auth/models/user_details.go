package models

import (
	"time"

	"gorm.io/datatypes"
)

// DB Model
type UserDetail struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	UserID         uint           `gorm:"uniqueIndex" json:"user_id"`
	Username       *string        `gorm:"uniqueIndex;size:50" json:"username"`
	Bio            string         `json:"bio"`
	ProfilePicture string         `json:"profile_picture"`
	SocialLinks    datatypes.JSON `json:"social_links"` // Matches JSONB column in DB
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

// Struct used for input binding
type UpdateUserDetailInput struct {
	Username       *string            `json:"username"`
	Bio            *string            `json:"bio"`
	ProfilePicture *string            `json:"profile_picture"`
	SocialLinks    *map[string]string `json:"social_links"`
}
