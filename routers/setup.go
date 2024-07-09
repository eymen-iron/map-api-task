package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupRouters(app *fiber.App, db *gorm.DB) {
	DB = db

	app.Get("/locations", GetLocation)
	app.Get("/location/:id", GetLocationByID)
	app.Post("/location/add", AddLocation)
	app.Put("/location/:id", UpdateLocationByID)
	app.Get("/locations/route", RouteLocation)

}
