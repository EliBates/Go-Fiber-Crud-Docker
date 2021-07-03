package main

import (
	property "github.com/elibates/go-fiber-example1/model"
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Delete("/property", property.DeleteProperty)
	app.Get("/property", property.GetProperty)
	app.Post("/property", property.CreateProperty)
	app.Get("/properties", property.GetProperties)
}

func main() {
	property.InitialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":8080")
}
