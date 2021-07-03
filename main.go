package main

import (
	property "github.com/elibates/go-fiber-example1/model"
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Post("/property", property.CreateProperty)
	app.Get("/property", property.GetProperty)
	app.Get("/properties", property.GetProperties)
	app.Put("/property", property.UpdateProperty)
	app.Delete("/property", property.DeleteProperty)
}

func main() {
	property.InitialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":8080")
}
