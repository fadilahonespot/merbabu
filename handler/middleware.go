package handler

import (
	"merbabu/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(logger.New(utils.LoggerConfig()))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, OPTIONS",
		AllowHeaders:     "Origin, Authorization, Access-Control-Allow-Origin, token, Content-Type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin",
		AllowCredentials: true,
	}))
}

