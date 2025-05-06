package models

import "time"

type ContentItem struct {
	ID            uint      `gorm:"primaryKey"`
	ContentTypeID uint      `gorm:"not null;index"`
	Title         string    `gorm:"size:100;not null"`
	Body          string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

// InputBody

type ContentItemRequest struct {
	ContentTypeID uint   `json:"content_type_id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Body          string `json:"body"`
}