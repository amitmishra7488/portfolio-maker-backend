package models


import (
	"time"
)

type Address struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `gorm:"index;not null" json:"user_id"`     
	Type       string    `gorm:"size:50;not null" json:"type"`      
	Line1      string    `gorm:"size:255;not null" json:"line1"`
	Line2      string    `gorm:"size:255" json:"line2,omitempty"`   
	City       string    `gorm:"size:100;not null" json:"city"`
	State      string    `gorm:"size:100;not null" json:"state"`
	Country    string    `gorm:"size:100;not null" json:"country"`
	Zipcode    string    `gorm:"size:20;not null" json:"zipcode"`
	IsDefault  bool      `gorm:"default:false" json:"is_default"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}


type AddressInput struct {
	Type      *string `json:"type"`
	Line1     *string `json:"line1"`
	Line2     *string `json:"line2"`
	City      *string `json:"city"`
	State     *string `json:"state"`
	Country   *string `json:"country"`
	Zipcode   *string `json:"zipcode"`
	IsDefault *bool   `json:"is_default"`
}

