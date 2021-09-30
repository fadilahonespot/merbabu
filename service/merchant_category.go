package service

import (
	"merbabu/constant"
	"merbabu/entity"
	"merbabu/utils"
)

func (s *defaultService) InsertMerchantCategory(req *entity.MerchantCategoryRequest) (err error) {
	_, err = s.merchantRepo.GetMerchantTypeById(req.MerchantTypeId)
	if err != nil {
		return
	}
	_, err = s.GetMerchantById(req.MerchantID)
	if err != nil {
		return
	}
	reqMerchant := entity.MerchantCategory{
		MerchantTypeId:      req.MerchantTypeId,
		Name:                req.Name,
		NeedDeliveryService: req.NeedDeliveryService,
	}

	tx := s.merchantRepo.BeginTrans()
	err = s.merchantRepo.InsertMerchantCategory(&reqMerchant, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	reqCategories := entity.MerchantCategories{
		MerchantID:       req.MerchantID,
		MerchantCategory: reqMerchant.IdMerchantCategory,
	}
	err = s.merchantRepo.InsertMerchantCategories(&reqCategories, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (s *defaultService) GetMerchantCatagoryById(id string) (resp *entity.MerchantCategory, err error) {
	return s.merchantRepo.GetMerchantCategoryById(id)
}

func (s *defaultService) UpdateMerchantCategory(categoryId, merchantId string, req *entity.MerchantCategoryUpdate) (err error) {
	merchantCategories, err := s.merchantRepo.GetMerchantCategorieByMerchantCategoryId(categoryId)
	if err != nil {
		return
	}
	if merchantCategories.MerchantID != merchantId {
		err = utils.SetError(constant.BadRequestCode, "merchantId is wrong")
		return
	}
	cat, err := s.merchantRepo.GetMerchantCategoryById(categoryId)
	if err != nil {
		return
	}
	_, err = s.merchantRepo.GetMerchantTypeById(req.MerchantTypeId)
	if err != nil {
		return
	}

	merchantCat := entity.MerchantCategory{
		IdMerchantCategory:  categoryId,
		MerchantTypeId:      req.MerchantTypeId,
		Name:                req.Name,
		NeedDeliveryService: req.NeedDeliveryService,
		CreateDate:          cat.CreateDate,
	}
	err = s.merchantRepo.UpdateMerchantCategory(&merchantCat)
	return
}

func (s *defaultService) GetMerchantCategoryList(page, length int, merchantId string) (resp *[]entity.MerchantCategoryList, err error) {
	_, err = s.merchantRepo.GetMerchantById(merchantId)
	if err != nil {
		return
	}
	return s.merchantRepo.GetMerchantCatagoryList(page, length, merchantId)
}

func (s *defaultService) DeleteMerchantCategory(categoryId, merchantId string) (err error) {
	tx := s.merchantRepo.BeginTrans()
	merchantCategories, err := s.merchantRepo.GetMerchantCategorieByMerchantCategoryId(categoryId)
	if err != nil {
		return
	}
	if merchantCategories.MerchantID != merchantId {
		err = utils.SetError(constant.BadRequestCode, "merchantId is wrong")
		return
	}
	err = s.merchantRepo.DeleteMerchantCategoriesByMerchantCategoryId(categoryId, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	err = s.merchantRepo.DeleteMerchantCategory(categoryId, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
