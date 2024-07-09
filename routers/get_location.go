package routers

import (
	"fmt"
	"strconv"

	"github.com/eymen-iron/map-api-task/models"
	"github.com/gofiber/fiber/v2"
)

func GetLocation(c *fiber.Ctx) error {
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

	var locations []models.Location
	if err := DB.Offset(offset).Limit(limit).Find(&locations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Location found successfully",
		"data":    locations,
	})
}

func GetLocationByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid location ID"})
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10000000
	}

	offset := (page - 1) * limit

	var location models.Location
	if err := DB.Offset(offset).Limit(limit).First(&location, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Database Error: %s", err.Error()),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Location found successfully",
		"data":    location,
	})

}
