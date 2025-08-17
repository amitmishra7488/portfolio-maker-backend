package content

import (
	"portfolio-user-service/repository/content"
	"portfolio-user-service/repository/content/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ContentItemService struct {
	Repo *content.ContentItemRepository
	Log  *zap.Logger
	DB   *gorm.DB
}

// Constructor
func NewContentItemService(repo *content.ContentItemRepository, db *gorm.DB, logger *zap.Logger) *ContentItemService {
	return &ContentItemService{
		Repo: repo,
		Log:  logger,
		DB:   db,
	}
}

func (s *ContentItemService) CreateContentItem(item *models.ContentItem) error {
	return s.Repo.Create(item)
}
