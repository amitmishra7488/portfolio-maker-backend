// service/address_service.go

package address

import (
	"errors"
	"portfolio-user-service/config"
	"portfolio-user-service/repository/address"
	"portfolio-user-service/repository/address/models"

	"gorm.io/gorm"
)

type AddressService struct{}

var (
	addressRepo = new(address.AddressRepository)
)

func (s AddressService) CreateAddress(userID uint, input models.AddressInput) (*models.Address, error) {
	var newAddress *models.Address

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Check if user already has any addresses
		hasAddress, err := addressRepo.UserHasAddressTx(tx, userID)
		if err != nil {
			return err
		}

		// Validate required fields
		if input.Type == nil || input.Line1 == nil || input.City == nil || input.State == nil || input.Country == nil || input.Zipcode == nil {
			return errors.New("missing required fields")
		}

		// Prepare new address object
		newAddress = &models.Address{
			UserID:    userID,
			Type:      *input.Type,
			Line1:     *input.Line1,
			Line2:     safeString(input.Line2),
			City:      *input.City,
			State:     *input.State,
			Country:   *input.Country,
			Zipcode:   *input.Zipcode,
			IsDefault: !hasAddress,
		}

		// Save the address
		if err := addressRepo.CreateAddressTx(tx, newAddress); err != nil {
			return err
		}

		return nil
	})

	return newAddress, err
}


func (s AddressService) UpdateAddress(userID, addressID uint, input models.AddressInput) (*models.Address, error) {
	// Get the address and ensure it belongs to the user
	address, err := addressRepo.GetAddressByIDAndUserID(config.DB, addressID, userID)
	if err != nil {
		return nil, err
	}

	// Build a map of fields to update (PATCH style)
	updates := make(map[string]interface{})

	if input.Type != nil {
		updates["type"] = *input.Type
	}
	if input.Line1 != nil {
		updates["line1"] = *input.Line1
	}
	if input.Line2 != nil {
		updates["line2"] = *input.Line2
	}
	if input.City != nil {
		updates["city"] = *input.City
	}
	if input.State != nil {
		updates["state"] = *input.State
	}
	if input.Country != nil {
		updates["country"] = *input.Country
	}
	if input.Zipcode != nil {
		updates["zipcode"] = *input.Zipcode
	}
	if input.IsDefault != nil {
		updates["is_default"] = *input.IsDefault
	}

	// Short-circuit if nothing to update
	if len(updates) == 0 {
		return address, nil // or return error if you want to enforce "must change something"
	}

	// Call repo with patch-style update
	if err := addressRepo.UpdateAddressFields(config.DB, addressID, userID, updates); err != nil {
		return nil, err
	}

	// Update local object with new values (for response)
	for k, v := range updates {
		switch k {
		case "type":
			address.Type = v.(string)
		case "line1":
			address.Line1 = v.(string)
		case "line2":
			address.Line2 = v.(string)
		case "city":
			address.City = v.(string)
		case "state":
			address.State = v.(string)
		case "country":
			address.Country = v.(string)
		case "zipcode":
			address.Zipcode = v.(string)
		case "is_default":
			address.IsDefault = v.(bool)
		}
	}

	return address, nil
}





func (s AddressService) GetAllAddresses(userID uint) ([]models.Address, error) {
	return addressRepo.GetAddressesByUserID(config.DB, userID)
}


func safeString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
