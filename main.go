package main

import (
	"log"
	"merbabu/handler"
	"merbabu/repository"
	"merbabu/service"
	"merbabu/utils/config"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	config.SetConfigServer()
	app := fiber.New(config.SetConfigFiber())
	db := config.SetDatabase()

	merchantRepo := repository.New().SetGormDB(db).Validate()
	merchantService := service.New().SetMerchantRepo(merchantRepo).Validate()
	merchantHandler := handler.New().SetMerchantService(merchantService).Validate()

	handler.SetRouterApp(app, merchantHandler)

	err := app.Listen(":" + viper.GetString("port"))
	if err != nil {
		log.Fatal("Failed Starting Server: ", err)
	}
}
