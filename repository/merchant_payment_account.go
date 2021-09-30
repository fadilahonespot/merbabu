package repository

import (
	"merbabu/entity"
	"merbabu/utils"
)

func (r *defaultRepository) InsertMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error) {
	err = r.db.Create(&req).Error
	return
}

func (r *defaultRepository) GetMerchantPaymentAccountList(page, length int, merchantId string) (resp *[]entity.MerchantPaymentAccountListView, err error) {
	err = r.db.Scopes(utils.Paginate(page, length)).Table("merchant_payment_account m").
		Select("m.id as payment_account_id, m.account_number, a.provider_name, a.type").
		Joins("INNER JOIN merchant_account_type a ON m.merchant_account_type = a.id").
		Where("m.merchant_id = ?", merchantId).Find(&resp).Error
	return
}

func (r *defaultRepository) GetMerchantPaymentAccountById(id string) (resp *entity.MerchantPaymentAccount, err error) {
	err = r.db.Where("id = ?", id).Take(&resp).Error
	return
}

func (r *defaultRepository) DeleteMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error) {
	err = r.db.Delete(&req).Error
	return
}

func (r *defaultRepository) UpdateMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error) {
	err = r.db.Save(&req).Error
	return
}

