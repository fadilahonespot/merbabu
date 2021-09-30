package entity

import (
	"time"

	uid "github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantCategory struct {
	IdMerchantCategory  string `json:"idMerchantCategory"`
	MerchantTypeId      string `json:"merchantTypeId"`
	Name                string `json:"name"`
	CreateDate          string `json:"createDate" gorm:"column:createDate"`
	NeedDeliveryService string `json:"needDeliveryService"`
}

type MerchantCategoryUpdate struct {
	MerchantTypeId      string `json:"merchantTypeId"`
	Name                string `json:"name"`
	NeedDeliveryService string `json:"needDeliveryService" validate:"required,eq=YES|eq=NO"`
}

type MerchantCategories struct {
	MerchantID       string `json:"merchatId"`
	MerchantCategory string `json:"merchantCategory"`
}

type MerchantCategoryRequest struct {
	MerchantID          string `json:"merchantId"`
	MerchantTypeId      string `json:"merchantTypeId"`
	Name                string `json:"name"`
	NeedDeliveryService string `json:"needDeliveryService" validate:"required,eq=YES|eq=NO"`
}

type MerchantCategoryList struct {
	IdMerchantCategory  string `json:"idMerchantCategory"`
	CategoryName        string `json:"categoryName"`
	NeedDeliveryService string `json:"needDeliveryService"`
	MerchantTypeId      string `json:"merchantTypeId"`
	TypeName            string `json:"typeName"`
	CreateDate          string `json:"create_date"`
}

func (e *MerchantCategory) TableName() string {
	return "merchant_category"
}

func (e *MerchantCategories) TableName() string {
	return "merchant_categories"
}

func (e *MerchantCategory) BeforeCreate(scope *gorm.DB) (err error) {
	e.IdMerchantCategory = uid.New().String()
	e.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	return
}
