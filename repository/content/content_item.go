package content

import (
	"portfolio-user-service/config"
	"portfolio-user-service/repository/content/models"
)

type ContentItemRepository struct{}

func (r *ContentItemRepository) Create(item *models.ContentItem) error {
	return config.DB.Create(item).Error
}
