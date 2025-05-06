package models

type ContentType struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null;index"`                              // Owner of this content type
	Name        string  `gorm:"size:100;not null;uniqueIndex:idx_user_name"` // Internal name (slug-like)
	Label       string  `gorm:"size:100;not null"`                           // Display label
	Description *string `gorm:"size:255"`                                    // Optional description
}

type ContentTypeInput struct {
	Name        *string `json:"name" binding:"required"`  // e.g., "painting"
	Label       *string `json:"label" binding:"required"` // e.g., "Painting"
	Description *string `json:"description,omitempty"`    // optional
}

// =========================================== RESPONSE================================

type ContentTypeWithItemsResponse struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Label       string                 `json:"label"`
	Description *string                `json:"description,omitempty"`
	Items       []ContentItemResponse  `json:"items"`
}

type ContentItemResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

