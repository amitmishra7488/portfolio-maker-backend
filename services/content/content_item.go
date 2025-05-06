package content

import (
	"portfolio-user-service/repository/content"
	"portfolio-user-service/repository/content/models"
)

type ContentItemService struct{}

var (
	contentItemRepo = new(content.ContentItemRepository)
)

func (s ContentItemService) CreateContentItem(item *models.ContentItem) error {
	return contentItemRepo.Create(item)
}
