package main

import (
	"github.com/atharv-bhadange/log-ingestor/db"
	"github.com/atharv-bhadange/log-ingestor/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := db.Connect()

	if err != nil {
		panic(err)
	}

	routes.Setup(app)

	app.Listen("localhost:3000")
}
