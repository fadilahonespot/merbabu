package repository

import (
	"merbabu/entity"
	"merbabu/utils"

	"gorm.io/gorm"
)

type defaultRepository struct {
	db *gorm.DB
}

func New() *defaultRepository {
	return &defaultRepository{}
}

func (r *defaultRepository) SetGormDB(db *gorm.DB) *defaultRepository {
	r.db = db
	return r
}

func (r *defaultRepository) Validate() MerchantRepo {
	if r.db == nil {
		panic("Database is nil")
	}
	return r
}

func (r *defaultRepository) InsertMerchant(req *entity.Merchant) (err error) {
	err = r.db.Create(&req).Error
	return
}

func (r *defaultRepository) GetMerchantById(id string) (resp *entity.Merchant, err error) {
	err = r.db.Where("id = ?", id).Take(&resp).Error
	return
}

func (r *defaultRepository) UpdateMerchant(req *entity.Merchant) (err error) {
	err = r.db.Save(&req).Error
	return
}

func (r *defaultRepository) DeleteMerchant(req *entity.Merchant) (err error) {
	err = r.db.Delete(&req).Error
	return
}

func (r *defaultRepository) GetMerchantList(page, length int) (resp *[]entity.Merchant, err error) {
	err = r.db.Scopes(utils.Paginate(page, length)).Find(&resp).Error
	return
}