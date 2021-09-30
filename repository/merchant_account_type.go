package repository

import (
	"merbabu/entity"
	"merbabu/utils"
)

func (r *defaultRepository) InsertMerchantAccountType(req *entity.MerchantAccountType) (err error) {
	err = r.db.Create(&req).Error
	return
}

func (r *defaultRepository) GetMerchantAccountTypeList(page, length int) (resp *[]entity.MerchantAccountType, err error) {
	err = r.db.Scopes(utils.Paginate(page, length)).Find(&resp).Error
	return
}

func (r *defaultRepository) GetMerchatAccountTypeById(id string) (resp *entity.MerchantAccountType, err error) {
	err = r.db.Where("id = ?", id).Take(&resp).Error
	return
}

func (r *defaultRepository) UpdateMerchantAccountType(req *entity.MerchantAccountType) (err error) {
	err = r.db.Save(&req).Error
	return
}