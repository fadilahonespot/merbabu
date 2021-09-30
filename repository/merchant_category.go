package repository

import (
	"merbabu/entity"
	"merbabu/utils"

	"gorm.io/gorm"
)

func (r *defaultRepository) InsertMerchantCategory(req *entity.MerchantCategory, tx *gorm.DB) (err error) {
	err = tx.Create(&req).Error
	return
}

func (r *defaultRepository) GetMerchantCategoryById(id string) (resp *entity.MerchantCategory, err error) {
	err = r.db.Where("id_merchant_category = ?", id).Take(&resp).Error
	return
}

func (r *defaultRepository) UpdateMerchantCategory(req *entity.MerchantCategory) (err error) {
	err = r.db.Where("id_merchant_category = ?", req.IdMerchantCategory).Save(&req).Error
	return
}

func (r *defaultRepository) GetMerchantCatagoryList(page, length int, merchantId string) (resp *[]entity.MerchantCategoryList, err error) {
	err = r.db.Scopes(utils.Paginate(page, length)).Table("merchant_category m").
		Select("m.id_merchant_category, m.name as category_name, m.need_delivery_service, t.id as merchant_type_id, t.name as type_name, m.createDate as create_date").
		Joins("JOIN merchant_categories c ON m.id_merchant_category = c.merchant_category JOIN merchant_type t ON t.id = m.merchant_type_id").
		Where("c.merchant_id = ?", merchantId).Find(&resp).Error
	return
}

func (r *defaultRepository) InsertMerchantCategories(req *entity.MerchantCategories, tx *gorm.DB) (err error) {
	err = tx.Create(&req).Error
	return
}

func (r *defaultRepository) DeleteMerchantCategoriesByMerchantCategoryId(id string, tx *gorm.DB) (err error) {
	err = tx.Where("merchant_category = ?", id).Delete(&entity.MerchantCategories{}).Error
	return
}

func (r *defaultRepository) DeleteMerchantCategory(id string, tx *gorm.DB) (err error) {
	err = tx.Where("id_merchant_category = ?", id).Delete(&entity.MerchantCategory{}).Error
	return
}

func (r *defaultRepository) GetMerchantCategorieByMerchantCategoryId(id string) (resp *entity.MerchantCategories, err error) {
	err = r.db.Where("merchant_category = ?", id).Take(&resp).Error
	return
}
