package main

import
( 
"github.com/gofiber/fiber/v2"
"github.com/Hamid-Afify9/Smart-API-With-Go-Docker-and-Postgres/database"
)

func main() {
    database.ConnectDb()
	app := fiber.New()
    setupRoutes(app)

	

	app.Listen(":3000")
}
