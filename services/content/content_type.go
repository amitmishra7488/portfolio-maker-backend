package content

import (
	"errors"
	content "portfolio-user-service/repository/content"
	"portfolio-user-service/repository/content/models"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ContentTypeService struct with dependencies
type ContentTypeService struct {
	Repo *content.ContentTypeRepository
	Log  *zap.Logger
	DB   *gorm.DB
}

// Constructor
func NewContentTypeService(repo *content.ContentTypeRepository, db *gorm.DB, logger *zap.Logger) *ContentTypeService {
	return &ContentTypeService{
		Repo: repo,
		Log:  logger,
		DB:   db,
	}
}

func (s *ContentTypeService) Create(userID uint, input models.ContentTypeInput) (*models.ContentType, error) {
	// Enforce max 3 content types per user
	count, err := s.Repo.CountByUser(userID)
	if err != nil {
		return nil, err
	}
	if count >= 3 {
		return nil, errors.New("you can only create up to 3 content types")
	}

	// Check for duplicate name
	exists, err := s.Repo.ExistsByName(userID, *input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("content type with this name already exists")
	}

	newType := &models.ContentType{
		UserID: userID,
		Name:   strings.ToLower(*input.Name),
		Label:  *input.Label,
	}

	if input.Description != nil {
		newType.Description = input.Description
	}

	if err := s.Repo.Create(newType); err != nil {
		return nil, err
	}

	return newType, nil
}

func (s *ContentTypeService) GetAllContentTypes(userID uint) ([]models.ContentTypeWithItemsResponse, error) {
	var res []models.ContentTypeWithItemsResponse
	res, err := s.Repo.GetAllContentTypes(userID)
	if err != nil {
		return res, err
	}
	return res, nil
}
