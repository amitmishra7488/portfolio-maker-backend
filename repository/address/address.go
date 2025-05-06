package address

import (
	"portfolio-user-service/repository/address/models"

	"gorm.io/gorm"
)

type AddressRepository struct{}

func (r *AddressRepository) UserHasAddressTx(tx *gorm.DB, userID uint) (bool, error) {
	var count int64
	if err := tx.Model(&models.Address{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}


// CreateAddressTx inserts a new address inside a transaction
func (r *AddressRepository) CreateAddressTx(tx *gorm.DB, address *models.Address) error {
	return tx.Create(address).Error
}

func (r *AddressRepository) GetAddressByIDAndUserID(db *gorm.DB, addressID, userID uint) (*models.Address, error) {
	var address models.Address
	err := db.Where("id = ? AND user_id = ?", addressID, userID).First(&address).Error
	return &address, err
}

func (r *AddressRepository) UpdateAddressFields(db *gorm.DB, addressID, userID uint, updates map[string]interface{}) error {
	return db.Model(&models.Address{}).
		Where("id = ? AND user_id = ?", addressID, userID).
		Updates(updates).Error
}

func (r *AddressRepository) GetAddressesByUserID(db *gorm.DB, userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := db.Where("user_id = ?", userID).Order("is_default DESC, created_at DESC").Find(&addresses).Error
	return addresses, err
}


