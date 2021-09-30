package repository

import (
	"merbabu/entity"
	"merbabu/utils"
)

func (r *defaultRepository) InsertMerchantType(req *entity.MerchantType) (err error) {
	err = r.db.Create(&req).Error
	return
}

func (r *defaultRepository) GetMerchantTypeById(id string) (resp *entity.MerchantType, err error) {
	err = r.db.Where("id = ?", id).Take(&resp).Error
	return
}

func (r *defaultRepository) UpdateMerchantType(req *entity.MerchantType) (err error) {
	err = r.db.Save(&req).Error
	return
}

func (r *defaultRepository) DeleteMerchantType(req *entity.MerchantType) (err error) {
	err = r.db.Delete(&req).Error
	return
}

func (r *defaultRepository) GetMerchantTypeList(page, length int) (resp *[]entity.MerchantType, err error) {
	err = r.db.Scopes(utils.Paginate(page, length)).Find(&resp).Error
	return
}
