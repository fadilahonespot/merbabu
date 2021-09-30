package config

import (
	"merbabu/constant"
	"merbabu/entity"
	"merbabu/utils"

	"github.com/gofiber/fiber/v2"
)

func SetConfigFiber() fiber.Config {
	conn := fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := utils.GetErrorCode(err)
			if code == "" {
				code = constant.GeneralErrorCode
			}

			resp := entity.DefaultResponse{
				Response: entity.Response{
					Status:  code,
					Message: err.Error(),
				},
				Data: struct{}{},
			}

			return c.JSON(resp)
		},
	}

	return conn
}
