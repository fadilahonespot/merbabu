package service

import (
	"merbabu/entity"
)

type MerchantService interface {
	GetMerchantById(id string) (resp *entity.Merchant, err error)
	InsertMerchant(req *entity.Merchant) (err error)
	UpdateMerchant(id string, req *entity.Merchant) (err error)
	DeleteMerchantById(id string) (err error)
	GetMerchantList(page, length int) (resp *[]entity.Merchant, err error)

	InsertMerchantType(req *entity.MerchantType) (err error)
	GetMerchantTypeById(id string) (resp *entity.MerchantType, err error)
	UpdateMerchantType(id string, req *entity.MerchantType) (err error) 
	DeleteMerchantTypeById(id string) (err error)
	GetMerchantTypeList(page, length int) (resp *[]entity.MerchantType, err error)

	InsertMerchantCategory(req *entity.MerchantCategoryRequest) (err error)
	GetMerchantCatagoryById(id string) (resp *entity.MerchantCategory ,err error)
	UpdateMerchantCategory(categoryId, merchantId string, req *entity.MerchantCategoryUpdate) (err error)
	GetMerchantCategoryList(page, length int, merchantId string) (resp *[]entity.MerchantCategoryList, err error)
	DeleteMerchantCategory(categoryId, merchantId string) (err error)

	InsertMerchantAccountType(req *entity.MerchantAccountType) (err error)
	GetMerchantAccountTypeList(page, length int) (resp *[]entity.MerchantAccountType, err error)
	UpdateMerchantAccountType(req *entity.MerchantAccountType, id string) (err error)

	InsertMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error)
	GetMerchantPaymentAccountList(page, length int, merchantId string) (resp *[]entity.MerchantPaymentAccountListView, err error)
	DeleteMerchantPaymentAccount(id string) (err error)
	UpdateMerchantPaymentAccount(idMerchantPaymentAccount string, req *entity.MerchantPaymentAccount) (err error)

	InsertOperatinalHours(req *entity.OperationalHours) (err error)
	GetOperatinalHoursByMerchatId(id string, page, length int) (resp *[]entity.OperationalHours, err error)
	UpdateOperatinalHours(operatinalHoursId, merchantId string, req *entity.OperationalHoursUpdate) (err error)
	DeleteOperatinalHours(operatinalHoursId, merchantId string) (err error)
}