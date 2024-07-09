package main

import (
	"log"

	"github.com/eymen-iron/map-api-task/models"
	"github.com/eymen-iron/map-api-task/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := models.GetDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	if err := db.AutoMigrate(&models.Location{}); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	routers.SetupRouters(app, db)

	err = app.Listen(":1337")
	if err != nil {
		log.Fatal(err)
	}
}
