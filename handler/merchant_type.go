package handler

import (
	"merbabu/entity"
	"merbabu/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *defaultHandler) GetMerchantType(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	resp, err := h.merchantService.GetMerchantTypeById(id)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}

func (h *defaultHandler) InsertMerchantType(c *fiber.Ctx) (err error) {
	var merchant entity.MerchantType
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.InsertMerchantType(&merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) UpdateMerchantType(c *fiber.Ctx) (err error) {
	var merchant entity.MerchantType
	id := c.Params("id")
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.UpdateMerchantType(id, &merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) DeleteMerchantType(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	err = h.merchantService.DeleteMerchantTypeById(id)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, nil)
}

func (h *defaultHandler) GetMerchantTypeList(c *fiber.Ctx) (err error) {
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

	resp, err := h.merchantService.GetMerchantTypeList(page, length)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}
