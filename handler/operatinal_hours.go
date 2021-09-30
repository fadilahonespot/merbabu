package handler

import (
	"merbabu/entity"
	"merbabu/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *defaultHandler) GetOperatinalHoursByMerhatId(c *fiber.Ctx) (err error) {
	id := c.Params("merchantId")
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
	resp, err := h.merchantService.GetOperatinalHoursByMerchatId(id, page, length)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, resp)
}

func (h *defaultHandler) InsertOperatinalHours(c *fiber.Ctx) (err error) {
	var merchant entity.OperationalHours
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.InsertOperatinalHours(&merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) UpdateOperationalHours(c *fiber.Ctx) (err error) {
	var merchant entity.OperationalHoursUpdate
	id := c.Params("id")
	merchatId := c.Params("merchantId")
	err = c.BodyParser(&merchant)
	if err != nil {
		return
	}
	err = utils.ValidateStruct(merchant)
	if err != nil {
		return
	}
	err = h.merchantService.UpdateOperatinalHours(id, merchatId, &merchant)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, merchant)
}

func (h *defaultHandler) DeleteOperatinalHours(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	merchatId := c.Params("merchantId")
	err = h.merchantService.DeleteOperatinalHours(id, merchatId)
	if err != nil {
		return
	}
	return utils.HandleSuccess(c, nil)
}

