package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-rest-api-logs/internal/interfaces"
	"go-rest-api-logs/internal/models"
)

// @Summary Crea un usuario en el sistema
// @Description Crea un usuario en el sistema.
// @Tags Administración
// @ID usuarios-create
// @Accept  json
// @Produce  json
// @Param param body models.UsuariosSysCreate true "Crear usuario"
// @Success 200 {array} models.UsuariosSysResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /users [post]
func CreateUsers(c *fiber.Ctx) error {
	users := new(models.UsuariosSysCreate)
	if err := c.BodyParser(&users); err != nil {
		return c.Status(503).JSON(models.HTTPError400{Code: 503, Message: "error"})
	}

	// comprobación de las contraseñas
	if users.Password != users.NewPassword {
		return c.Status(fiber.StatusOK).JSON(models.HTTPError400{Code: 400, Message: "Las contraseñas no coinciden."})
	}

	// creando el objeto para agregar al usuario
	var usersSys = models.UsuariosSys{Identificador: uuid.New().String(), Login: users.Login, Password: users.Password, IsAdmin: false}
	var resultUsers, err, result = interfaces.CreateUsers(usersSys)

	if result {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON(fiber.Map{"identificador": resultUsers.Identificador, "login": resultUsers.Login, "is_admin": resultUsers.IsAdmin})
}

// @Summary Eliminar un usuario del sistema
// @Description Elimina el usuario del sistema.
// @Tags Administración
// @ID usuarios-delete
// @Accept  json
// @Produce  json
// @Success 200 {array} models.UsuariosSysResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /users [delete]
func UsuariosSysDelete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "get all",
	})
}
