package handler

import "github.com/gofiber/fiber/v2"

func SetRouterApp(app *fiber.App, handler *defaultHandler) {
	SetupMiddleware(app)

	app.Post("/merchant", handler.InsertMerchant)
	app.Get("/merchant/:id", handler.GetMerchant)
	app.Put("/merchant/:id", handler.UpdateMerchant)
	app.Delete("/merchant/:id", handler.DeleteMerchant)
	app.Get("/merchants", handler.GetMerchantList)

	app.Post("/type", handler.InsertMerchantType)
	app.Get("/type/:id", handler.GetMerchantType)
	app.Put("/type/:id", handler.UpdateMerchantType)
	app.Delete("/type/:id", handler.DeleteMerchantType)
	app.Get("/types", handler.GetMerchantTypeList)

	app.Post("/category", handler.InsertMerchantCategory)
	app.Get("/category/:merchantId", handler.GetMerchantCategoryList)
	app.Put("/category/:merchantId/:merchantCategoryId", handler.UpdateMerchantCategory)
	app.Delete("/category/:merchantId/:merchantCategoryId", handler.DeleteMerchantCategory)

	app.Post("/account-type", handler.InsertMerchantAccountType)
	app.Get("/account-types", handler.GetMerchantAccountTypeList)
	app.Put("/account-type/:id", handler.UpdateMerchantAccountType)
	
	app.Post("/account-payment", handler.InsertMerchantPaymentAccount)
	app.Get("/account-payment/:merchantId", handler.GetMerchantPaymentAccountList)
	app.Delete("/account-payment/:id", handler.DeleteMerchantPaymentAccount)
	app.Put("/account-payment/:id", handler.UpdateMerchantPaymentAccount)

	app.Post("/operatinal-hours", handler.InsertOperatinalHours)
	app.Get("/operatinal-hours/:merchantId", handler.GetOperatinalHoursByMerhatId)
	app.Put("/operatinal-hours/:merchantId/:id", handler.UpdateOperationalHours)
	app.Delete("/operatinal-hours/:merchantId/:id", handler.DeleteOperatinalHours)
}