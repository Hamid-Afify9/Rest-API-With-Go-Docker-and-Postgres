package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Hamid-Afify9/Smart-API-With-Go-Docker-and-Postgres/handlers"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListFacts)
	app.Post("/create", handlers.CreateFact)
}
