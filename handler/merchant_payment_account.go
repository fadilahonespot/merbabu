package handler

import (
	"merbabu/entity"
	"merbabu/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *defaultHandler) InsertMerchantPaymentAccount(c *fiber.Ctx) (err error) {
	var merchant entity.MerchantPaymentAccount
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.InsertMerchantPaymentAccount(&merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) GetMerchantPaymentAccountList(c *fiber.Ctx) (err error) {
	pageStr := c.Query("page")
	lengthStr := c.Query("length")
	merchantId := c.Params("merchantId")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return
	}

	resp, err := h.merchantService.GetMerchantPaymentAccountList(page, length, merchantId)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}

func (h *defaultHandler) DeleteMerchantPaymentAccount(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	err = h.merchantService.DeleteMerchantPaymentAccount(id)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, nil)
}

func (h *defaultHandler) UpdateMerchantPaymentAccount(c *fiber.Ctx) (err error) {
	var merchant entity.MerchantPaymentAccount
	id := c.Params("id")
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.UpdateMerchantPaymentAccount(id, &merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}


