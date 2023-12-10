package main

import (
	"github.com/bshome19/programming-languages-with-db/configs"
	"github.com/bshome19/programming-languages-with-db/routers"
	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Everyone! Let's build REST APIs in Go fiber using Gorm and SQLite.")
	})
	routers.RouterInstance().SetupRoutes(app)
	
	configs.DatabaseConnectionInstance().ConnectDB()

	app.Listen(":8080")
}
