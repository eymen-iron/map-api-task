package routers

import (
	"strconv"
	"strings"

	"github.com/eymen-iron/map-api-task/models"
	"github.com/eymen-iron/map-api-task/utils"
	"github.com/gofiber/fiber/v2"
)

func RouteLocation(c *fiber.Ctx) error {

	longStr := c.Query("long")
	latStr := c.Query("lat")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10000000
	}

	offset := (page - 1) * limit

	if longStr == "" || latStr == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Missing required parameters",
		})
	}

	latFloat, err := strconv.ParseFloat(strings.TrimSpace(latStr), 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Please fill the lat field correctly",
		})

	}
	longFloat, err := strconv.ParseFloat(strings.TrimSpace(longStr), 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Please fill the long field correctly",
		})
	}

	l := models.Location{
		Latitude:  latFloat,
		Longitude: longFloat,
	}

	if err := utils.Validate(l); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	locations := []models.Location{}

	err = DB.Raw("CALL GetNearestLocationsWithPagination(?, ?, ?, ?)", limit, offset, latFloat, longFloat).Scan(&locations).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "Location routed successfully",
		"data":    locations,
	})
}
