package content

import (
	"portfolio-user-service/config"
	"portfolio-user-service/repository/content/models"

	"gorm.io/gorm"
)

type ContentTypeRepository struct{}

func (r *ContentTypeRepository) Create(ct *models.ContentType) error {
	return config.DB.Create(ct).Error
}

func (r *ContentTypeRepository) CountByUser(userID uint) (int64, error) {
	var count int64
	err := config.DB.Model(&models.ContentType{}).
		Where("user_id = ?", userID).
		Count(&count).Error
	return count, err
}

func (r *ContentTypeRepository) ExistsByName(userID uint, name string) (bool, error) {
	var ct models.ContentType
	err := config.DB.
		Where("user_id = ? AND name = ?", userID, name).
		First(&ct).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func (r *ContentTypeRepository) GetAllContentTypes(userID uint) ([]models.ContentTypeWithItemsResponse, error) {
	var result []models.ContentTypeWithItemsResponse
	// Step 1: Fetch all content types for the user
	var contentTypes []models.ContentType
	if err := config.DB.Where("user_id = ?", userID).Find(&contentTypes).Error; err != nil {
		return result, err
	}

	// Step 2: Collect all content_type_ids
	var contentTypeIDs []uint
	for _, ct := range contentTypes {
		contentTypeIDs = append(contentTypeIDs, ct.ID)
	}

	// Step 3: Fetch all content items in one query
	var items []models.ContentItem
	if len(contentTypeIDs) > 0 {
		if err := config.DB.Where("content_type_id IN ?", contentTypeIDs).Find(&items).Error; err != nil {
			return result, err
		}
	}

	// Step 4: Group content items by content_type_id
	itemMap := make(map[uint][]models.ContentItemResponse)
	for _, item := range items {
		itemMap[item.ContentTypeID] = append(itemMap[item.ContentTypeID], models.ContentItemResponse{
			ID:    item.ID,
			Title: item.Title,
			Body:  item.Body,
		})
	}

	// Step 5: Assemble final response
	for _, ct := range contentTypes {
		result = append(result, models.ContentTypeWithItemsResponse{
			ID:          ct.ID,
			Name:        ct.Name,
			Label:       ct.Label,
			Description: ct.Description,
			Items:       itemMap[ct.ID],
		})
	}

	return result, nil
}
