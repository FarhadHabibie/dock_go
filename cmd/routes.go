package main

import (
	"github.com/FarhadHabibie/dock_go/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
	app.Get("/list", handlers.ListFacts)
	app.Post("/delfact", handlers.DeleteFacts)
	app.Post("/drpfact", handlers.DropFacts)
}
