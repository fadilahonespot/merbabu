package service

import "merbabu/entity"

func (s *defaultService) InsertMerchantAccountType(req *entity.MerchantAccountType) (err error) {
	return s.merchantRepo.InsertMerchantAccountType(req)
}

func (s *defaultService) GetMerchantAccountTypeList(page, length int) (resp *[]entity.MerchantAccountType, err error) {
	return s.merchantRepo.GetMerchantAccountTypeList(page, length)
}

func (s *defaultService) UpdateMerchantAccountType(req *entity.MerchantAccountType, id string) (err error) {
	_, err = s.merchantRepo.GetMerchatAccountTypeById(id)
	if err != nil {
		return
	}
	req.ID = id
	return s.merchantRepo.UpdateMerchantAccountType(req)
}