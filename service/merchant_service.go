package service

import (
	"merbabu/entity"
	"merbabu/repository"
)

type defaultService struct {
	merchantRepo repository.MerchantRepo
}

func New() *defaultService {
	return &defaultService{}
}

func (s *defaultService) SetMerchantRepo(r repository.MerchantRepo) *defaultService {
	s.merchantRepo = r
	return s
}

func (s *defaultService) Validate() MerchantService {
	if s.merchantRepo == nil {
		panic("merchantRepo is nil")
	}
	return s
}

func (s *defaultService) GetMerchantById(id string) (resp *entity.Merchant, err error) {
	return s.merchantRepo.GetMerchantById(id)
}

func (s *defaultService) InsertMerchant(req *entity.Merchant) (err error) {
	return s.merchantRepo.InsertMerchant(req)
}

func (s *defaultService) UpdateMerchant(id string, req *entity.Merchant) (err error) {
	_, err = s.merchantRepo.GetMerchantById(id)
	if err != nil {
		return
	}
	
	req.ID = id
	err = s.merchantRepo.UpdateMerchant(req)
	return
}

func (s *defaultService) DeleteMerchantById(id string) (err error) {
	var req entity.Merchant
	req.ID = id

	_, err = s.merchantRepo.GetMerchantById(id)
	if err != nil {
		return
	}
	err = s.merchantRepo.DeleteMerchant(&req)
	return
}

func (s *defaultService) GetMerchantList(page, length int) (resp *[]entity.Merchant, err error) {
	return s.merchantRepo.GetMerchantList(page, length)
}

