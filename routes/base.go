package routes

import (
	"github.com/atharv-bhadange/log-ingestor/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controller.HealthCheck)

	app.Post("/log", controller.AddLog)

	app.Get("/log", controller.GetLog)

	app.Get("/log/:id", controller.GetLogById)

	app.Get("/log/level/:level", controller.GetLogsByLevel)

}
