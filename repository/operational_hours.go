package repository

import (
	"merbabu/entity"
	"merbabu/utils"
)

func (r *defaultRepository) InsertOperationalHours(req *entity.OperationalHours) (err error) {
	err = r.db.Create(&req).Error
	return
}

func (r *defaultRepository) GetOperationalHoursById(id string) (resp *entity.OperationalHours, err error) {
	err = r.db.Where("id = ?", id).Take(&resp).Error
	return
}

func (r *defaultRepository) GetOperationalHoursByMerchantId(id string, page, length int) (resp *[]entity.OperationalHours, err error) {
	err = r.db.Scopes(utils.Paginate(page, length)).Where("merchant_id = ?", id).Find(&resp).Error
	return
}

func (r *defaultRepository) UpdateOperationalHours(req *entity.OperationalHours) (err error) {
	err = r.db.Save(&req).Error
	return
}

func (r *defaultRepository) DeleteOperationalHours(req *entity.OperationalHours) (err error) {
	err = r.db.Delete(&req).Error
	return
}