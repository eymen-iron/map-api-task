package routers

import (
	"strconv"
	"strings"

	"github.com/eymen-iron/map-api-task/models"
	"github.com/eymen-iron/map-api-task/utils"
	"github.com/gofiber/fiber/v2"
)

func AddLocation(c *fiber.Ctx) error {
	name := c.FormValue("name")
	lat := c.FormValue("lat")
	long := c.FormValue("long")
	marker := c.FormValue("marker")

	if name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Please fill the name field",
		})
	} else if lat == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Please fill the lat field",
		})
	} else if long == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Please fill the long field",
		})
	} else if marker == "" {
		marker = "red"
	}

	latFloat, err := strconv.ParseFloat(strings.TrimSpace(lat), 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Please fill the lat field correctly",
		})

	}

	longFloat, err := strconv.ParseFloat(strings.TrimSpace(long), 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Please fill the long field correctly",
		})
	}

	location := models.Location{
		Name:      name,
		Latitude:  latFloat,
		Longitude: longFloat,
		Marker:    marker,
	}

	if err := utils.Validate(location); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	result := DB.Create(&location)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": "Something went wrong",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "Location added successfully",
	})

}
