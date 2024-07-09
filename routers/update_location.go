package routers

import (
	"strconv"
	"strings"

	"github.com/eymen-iron/map-api-task/models"
	"github.com/eymen-iron/map-api-task/utils"
	"github.com/gofiber/fiber/v2"
)

func UpdateLocationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	name := c.FormValue("name")
	latStr := c.FormValue("lat")
	lngStr := c.FormValue("lng")
	marker := c.FormValue("marker")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "id is required",
		})
	}

	locationID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid id format",
		})
	}

	location := models.Location{ID: locationID}
	result := DB.First(&location)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "location not found",
		})
	}

	if name != "" {
		location.Name = name
	}

	if marker != "" {
		location.Marker = marker
	}

	if latStr != "" {
		latFloat, err := strconv.ParseFloat(strings.TrimSpace(latStr), 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":   true,
				"message": "Please fill the lat field correctly",
			})

		}
		location.Latitude = latFloat
	}

	if lngStr != "" {
		lngFloat, err := strconv.ParseFloat(strings.TrimSpace(lngStr), 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":   true,
				"message": "Please fill the lng field correctly",
			})

		}
		location.Longitude = lngFloat
	}

	if err := utils.Validate(location); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	DB.Save(&location)

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "location updated successfully",
	})
}
