package service

import "merbabu/entity"

func (s *defaultService) GetMerchantTypeById(id string) (resp *entity.MerchantType, err error) {
	return s.merchantRepo.GetMerchantTypeById(id)
}

func (s *defaultService) InsertMerchantType(req *entity.MerchantType) (err error) {
	return s.merchantRepo.InsertMerchantType(req)
}

func (s *defaultService) UpdateMerchantType(id string, req *entity.MerchantType) (err error) {
	_, err = s.merchantRepo.GetMerchantTypeById(id)
	if err != nil {
		return
	}

	req.ID = id
	err = s.merchantRepo.UpdateMerchantType(req)
	return
}

func (s *defaultService) DeleteMerchantTypeById(id string) (err error) {
	var req entity.MerchantType
	req.ID = id

	_, err = s.merchantRepo.GetMerchantTypeById(id)
	if err != nil {
		return
	}
	err = s.merchantRepo.DeleteMerchantType(&req)
	return
}

func (s *defaultService) GetMerchantTypeList(page, length int) (resp *[]entity.MerchantType, err error) {
	return s.merchantRepo.GetMerchantTypeList(page, length)
}
