package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go_admin.com/database"
	"go_admin.com/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
