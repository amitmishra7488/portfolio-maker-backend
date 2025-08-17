package content

import (
	"portfolio-user-service/repository/content/models"

	"gorm.io/gorm"
)

type ContentItemRepository struct {
	DB *gorm.DB
}

func NewContentItemRepository(db *gorm.DB) *ContentItemRepository {
	return &ContentItemRepository{
		DB: db,
	}
}

func (r *ContentItemRepository) Create(item *models.ContentItem) error {
	return r.DB.Create(item).Error
}
