package entity

import (
	"time"

	uid "github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantType struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (e *MerchantType) TableName() string {
	return "merchant_type"
}

func (e *MerchantType) BeforeCreate(scope *gorm.DB) (err error) {
	e.ID = uid.New().String()
	return
}
