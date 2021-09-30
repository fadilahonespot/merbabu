package service

import "merbabu/entity"

func (s *defaultService) InsertMerchantPaymentAccount(req *entity.MerchantPaymentAccount) (err error) {
	_, err = s.merchantRepo.GetMerchantById(req.MerchantID)
	if err != nil {
		return
	}
	_, err = s.merchantRepo.GetMerchatAccountTypeById(req.MerchantAccountType)
	if err != nil {
		return
	}
	err = s.merchantRepo.InsertMerchantPaymentAccount(req)
	return
}

func (s *defaultService) GetMerchantPaymentAccountList(page, length int, merchantId string) (resp *[]entity.MerchantPaymentAccountListView, err error) {
	_, err = s.merchantRepo.GetMerchantById(merchantId)
	if err != nil {
		return
	}
	return s.merchantRepo.GetMerchantPaymentAccountList(page, length, merchantId)
}

func (s *defaultService) DeleteMerchantPaymentAccount(id string) (err error) {
	_, err = s.merchantRepo.GetMerchantPaymentAccountById(id)
	if err != nil {
		return
	}
	var req entity.MerchantPaymentAccount
	req.ID = id
	return s.merchantRepo.DeleteMerchantPaymentAccount(&req)
}

func (s *defaultService) UpdateMerchantPaymentAccount(idMerchantPaymentAccount string, req *entity.MerchantPaymentAccount) (err error) {
	_, err = s.merchantRepo.GetMerchantPaymentAccountById(idMerchantPaymentAccount)
	if err != nil {
		return
	}
	_, err = s.merchantRepo.GetMerchantById(req.MerchantID)
	if err != nil {
		return
	}
	_, err = s.merchantRepo.GetMerchatAccountTypeById(req.MerchantAccountType)
	if err != nil {
		return
	}
	req.ID = idMerchantPaymentAccount
	err = s.merchantRepo.UpdateMerchantPaymentAccount(req)
	return
}