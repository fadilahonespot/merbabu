package handler

import (
	"merbabu/entity"
	"merbabu/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *defaultHandler) InsertMerchantAccountType(c *fiber.Ctx) (err error) {
	var merchant entity.MerchantAccountType
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.InsertMerchantAccountType(&merchant)
	if err != nil {
		return
	}

	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) GetMerchantAccountTypeList(c *fiber.Ctx) (err error) {
	pageStr := c.Query("page")
	lengthStr := c.Query("length")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return
	}

	resp, err := h.merchantService.GetMerchantAccountTypeList(page, length)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}

func (h *defaultHandler) UpdateMerchantAccountType(c *fiber.Ctx) (err error) {
	var merchant entity.MerchantAccountType
	id := c.Params("id")
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.UpdateMerchantAccountType(&merchant, id)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}
