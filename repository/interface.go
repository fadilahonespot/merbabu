package repository

import (
	"merbabu/entity"

	"gorm.io/gorm"
)

type MerchantRepo interface {
	//Transaction db
	BeginTrans() *gorm.DB

	// Merchant
	InsertMerchant(req *entity.Merchant) (err error)
	GetMerchantById(id string) (resp *entity.Merchant, err error) 
	UpdateMerchant(req *entity.Merchant) (err error)
	DeleteMerchant(req *entity.Merchant) (err error)
	GetMerchantList(page, length int) (resp *[]entity.Merchant, err error)

	// Merchant Type
	InsertMerchantType(req *entity.MerchantType) (err error)
	GetMerchantTypeById(id string) (resp *entity.MerchantType, err error)
	UpdateMerchantType(req *entity.MerchantType) (err error) 
	DeleteMerchantType(req *entity.MerchantType) (err error)
	GetMerchantTypeList(page, length int) (resp *[]entity.MerchantType, err error)

	// Merchant Category
	InsertMerchantCategory(req *entity.MerchantCategory, tx *gorm.DB) (err error)
	GetMerchantCategoryById(id string) (resp *entity.MerchantCategory, err error) 
	UpdateMerchantCategory(req *entity.MerchantCategory) (err error)
	GetMerchantCatagoryList(page, length int, merchantId string) (resp *[]entity.MerchantCategoryList, err error)
	InsertMerchantCategories(req *entity.MerchantCategories, tx *gorm.DB) (err error)
	DeleteMerchantCategory(id string, tx *gorm.DB) (err error)
	DeleteMerchantCategoriesByMerchantCategoryId(id string, tx *gorm.DB) (err error)
	GetMerchantCategorieByMerchantCategoryId(id string) (resp *entity.MerchantCategories, err error)

	// Merchant Account Type
	InsertMerchantAccountType(req *entity.MerchantAccountType) (err error)
	GetMerchantAccountTypeList(page, length int) (resp *[]entity.MerchantAccountType, err error)
	GetMerchatAccountTypeById(id string) (resp *entity.MerchantAccountType, err error)
	UpdateMerchantAccountType(req *entity.MerchantAccountType) (err error)

	// Merchant Payment Account
	InsertMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error)
	GetMerchantPaymentAccountList(page, length int, merchantId string) (resp *[]entity.MerchantPaymentAccountListView, err error)
	GetMerchantPaymentAccountById(id string) (resp *entity.MerchantPaymentAccount, err error)
	DeleteMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error)
	UpdateMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error)

	// Operatinal Hours
	InsertOperationalHours(req *entity.OperationalHours) (err error)
	GetOperationalHoursById(id string) (resp *entity.OperationalHours, err error)
	UpdateOperationalHours(req *entity.OperationalHours) (err error)
	DeleteOperationalHours(req *entity.OperationalHours) (err error)
	GetOperationalHoursByMerchantId(id string, page, length int) (resp *[]entity.OperationalHours, err error)
}