package handler

import (
	"merbabu/entity"
	"merbabu/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *defaultHandler) InsertMerchantCategory(c *fiber.Ctx) (err error) {
	var merchant entity.MerchantCategoryRequest
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.InsertMerchantCategory(&merchant)
	if err != nil {
		return
	}

	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) GetMerchantCategoryList(c *fiber.Ctx) (err error) {
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

	resp, err := h.merchantService.GetMerchantCategoryList(page, length, merchantId)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}

func (h *defaultHandler) DeleteMerchantCategory(c *fiber.Ctx) (err error) {
	merchantId := c.Params("merchantId")
	merchanCategoryId := c.Params("merchantCategoryId")
	
	err = h.merchantService.DeleteMerchantCategory(merchanCategoryId, merchantId)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, nil)
}

func (h *defaultHandler) UpdateMerchantCategory(c *fiber.Ctx) (err error) {
	merchantId := c.Params("merchantId")
	merchanCategoryId := c.Params("merchantCategoryId")

	var merchant entity.MerchantCategoryUpdate
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	
	err = h.merchantService.UpdateMerchantCategory(merchanCategoryId, merchantId, &merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}
