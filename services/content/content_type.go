package content

import (
	"errors"
	content "portfolio-user-service/repository/content"
	"portfolio-user-service/repository/content/models"
	"strings"
)

type ContentTypeService struct{}

var (
	contentTypeRepo = new(content.ContentTypeRepository)
)

func (s ContentTypeService) Create(userID uint, input models.ContentTypeInput) (*models.ContentType, error) {
	// Enforce max 3 content types per user
	count, err := contentTypeRepo.CountByUser(userID)
	if err != nil {
		return nil, err
	}
	if count >= 3 {
		return nil, errors.New("you can only create up to 3 content types")
	}

	// Check for duplicate name
	exists, err := contentTypeRepo.ExistsByName(userID, *input.Name)
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

	if err := contentTypeRepo.Create(newType); err != nil {
		return nil, err
	}

	return newType, nil
}

func (s ContentTypeService) GetAllContentTypes(userID uint) ([]models.ContentTypeWithItemsResponse, error) {
	var res []models.ContentTypeWithItemsResponse
	res, err := contentTypeRepo.GetAllContentTypes(userID)
	if err != nil {
		return res, err
	}
	return res, nil
}