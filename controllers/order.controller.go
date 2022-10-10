package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go_admin.com/database"
	"go_admin.com/models"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
