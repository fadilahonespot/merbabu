package entity

import (
	uid "github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantAccountType struct {
	ID           string `json:"id"`
	Type         string `json:"type" validate:"required,eq=BANK|eq=EWALLET"`
	ProviderName string `json:"providerName"`
}

func (e *MerchantAccountType) TableName() string {
	return "merchant_account_type"
}

func (e *MerchantAccountType) BeforeCreate(scope *gorm.DB) (err error) {
	e.ID = uid.New().String()
	return
}
