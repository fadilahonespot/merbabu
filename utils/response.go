package utils

import (
	"merbabu/entity"
	"merbabu/constant"

	"github.com/gofiber/fiber/v2"
)

func HandleSuccess(c *fiber.Ctx, data interface{}) error {
	resp := entity.DefaultResponse{
		Response: entity.Response{
			Status: constant.ResponseSucessCode,
			Message: "Success",
		},
		Data: data,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}