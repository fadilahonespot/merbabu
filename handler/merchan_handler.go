package handler

import (
	"merbabu/entity"
	"merbabu/service"
	"merbabu/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type defaultHandler struct {
	merchantService service.MerchantService
}

func New() *defaultHandler {
	return &defaultHandler{}
}

func (h *defaultHandler) SetMerchantService(r service.MerchantService) *defaultHandler {
	h.merchantService = r
	return h
}

func (h *defaultHandler) Validate() *defaultHandler {
	if h.merchantService == nil {
		panic("merchantService is nil")
	}
	return h
}

func (h *defaultHandler) GetMerchant(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	resp, err := h.merchantService.GetMerchantById(id)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}

func (h *defaultHandler) InsertMerchant(c *fiber.Ctx) (err error) {
	var merchant entity.Merchant
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.InsertMerchant(&merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) UpdateMerchant(c *fiber.Ctx) (err error) {
	var merchant entity.Merchant
	id := c.Params("id")
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.UpdateMerchant(id, &merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) DeleteMerchant(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	err = h.merchantService.DeleteMerchantById(id)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, nil)
}

func (h *defaultHandler) GetMerchantList(c *fiber.Ctx) (err error) {
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

	resp, err := h.merchantService.GetMerchantList(page, length)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}

