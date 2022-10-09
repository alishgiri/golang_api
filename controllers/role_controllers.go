package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go_admin.com/database"
	"go_admin.com/models"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map // DTO - Data Transfer Object

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	if len(list) == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Permission cannot be empty",
		})
	}

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permisssionId := range list {
		permId, _ := strconv.Atoi(permisssionId.(string))

		permissions[i] = models.Permission{
			Id: uint(permId),
		}
	}

	var result fiber.Map
	database.DB.Table("role_permissions").Delete(&result, "role_id = ?", id)
	// Or use below code
	// database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}
