package service

import (
	"merbabu/constant"
	"merbabu/entity"
	"merbabu/utils"
)

func (s *defaultService) InsertOperatinalHours(req *entity.OperationalHours) (err error) {
	_, err = s.merchantRepo.GetMerchantById(req.MerchantID)
	if err != nil {
		return
	}
	err = s.merchantRepo.InsertOperationalHours(req)
	return
}

func (s *defaultService) GetOperatinalHoursByMerchatId(id string, page, length int) (resp *[]entity.OperationalHours, err error) {
	_, err = s.merchantRepo.GetMerchantById(id)
	if err != nil {
		return
	}
	return s.merchantRepo.GetOperationalHoursByMerchantId(id, page, length)
}

func (s *defaultService) UpdateOperatinalHours(operatinalHoursId, merchantId string, req *entity.OperationalHoursUpdate) (err error) {
	opWork, err := s.merchantRepo.GetOperationalHoursById(operatinalHoursId)
	if err != nil {
		return
	}
	if opWork.MerchantID != merchantId {
		err = utils.SetError(constant.BadRequestCode, "merchatId is wrong")
		return
	}
	work := entity.OperationalHours{
		ID:        operatinalHoursId,
		MerchantID: merchantId,
		Day:       req.Day,
		OpenHour:  req.OpenHour,
		CloseHour: req.CloseHour,
	}
	err = s.merchantRepo.UpdateOperationalHours(&work)
	return
}

func (s *defaultService) DeleteOperatinalHours(operatinalHoursId, merchantId string) (err error) {
	opWork, err := s.merchantRepo.GetOperationalHoursById(operatinalHoursId)
	if err != nil {
		return
	}
	if opWork.MerchantID != merchantId {
		err = utils.SetError(constant.BadRequestCode, "merchatId is wrong")
		return
	}
	var opHours entity.OperationalHours
	opHours.ID = operatinalHoursId
	err = s.merchantRepo.DeleteOperationalHours(&opHours)
	return
}
