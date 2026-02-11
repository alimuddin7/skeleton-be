package routers

import (
	"test-new/configs"
	"test-new/controllers"
	"test-new/helpers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/rs/zerolog"
)

func InitializeRouter(ctrl controllers.Controller, l zerolog.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: configs.Cfg.General.ServiceName,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(helpers.TracingMiddleware())
	app.Use(helpers.LoggingMiddleware(l))

	main := app.Group(configs.Cfg.Server.MainPath)
	
	v1 := main.Group("/v1")
	v1.Get("/healthcheck", ctrl.V1().HealthCheck)
	// [V1_ROUTES_MARKER]

	return app
}
