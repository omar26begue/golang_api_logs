package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-rest-api-logs/internal/database"
	"go-rest-api-logs/internal/interfaces"
	"go-rest-api-logs/internal/models"
)

// @Summary Devuelve todas las OSDE
// @Description Devuelve un listado de todas las OSDES en el sistemas.
// @Tags OSDE
// @Accept  json
// @Produce  json
// @Success 200 {array} models.OsdeResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/osde [get]
func OsdeGetAll(c *fiber.Ctx) error {
	var osde = interfaces.GetAll()

	return c.Status(fiber.StatusOK).JSON(osde)
}

// @Summary Devuelve todas las OSDE
// @Description Devuelve un listado de todas las OSDES en el sistemas.
// @Tags OSDE
// @Accept  json
// @Produce  json
// @Success 200 {array} models.OsdeResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v2/osde [get]
func OsdeGetAllv2(c *fiber.Ctx) error {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var osde models.OsdeResponse
	db.Find(&osde)

	return c.Status(fiber.StatusOK).JSON(osde)
}

// @Summary Devuelve la OSDE enviada por parámetro
// @Description Devuelve la OSDE enviada por parámetro.
// @Tags OSDE
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "OSDE"
// @Success 200 {object} models.OsdeResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/osde/{id} [get]
func OsdeGetFindOne(c *fiber.Ctx) error {
	idosde := c.Params("id")

	result, count := interfaces.GetFindOneOSDE(idosde)

	if count == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&models.HTTPError400{Code: 404, Message: "No se encuentra la OSDE solicitada"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// @Summary Crea una OSDE
// @Description get string by ID
// @Tags OSDE
// @Accept  json
// @Produce  json
// @Param osde body models.OsdeRequest true "Adicionar OSDE"
// @Success 201 {object} models.OsdeResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/osde [post]
func OsdeCreate(c *fiber.Ctx) error {
	osde := new(models.Osde)
	if err := c.BodyParser(&osde); err != nil {
		return c.Status(503).JSON(models.HTTPError400{Code: 503, Message: "error"})
	}

	osde.Identificador = uuid.New().String()
	var resultOsde, err, result = interfaces.CreateOsde(osde)

	if result == true {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON(resultOsde)
}

// @Summary Modifica una OSDE
// @Description get string by ID
// @Tags OSDE
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "OSDE"
// @Param osde body models.OsdeRequest true "Actualizar datos OSDE"
// @Success 200 {object} models.OsdeResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/osde/{id} [patch]
func OsdeUpdate(c *fiber.Ctx) error {
	idosde := c.Params("id")

	osde := new(models.Osde)
	if err := c.BodyParser(&osde); err != nil {
		return c.Status(503).JSON(models.HTTPError400{Code: 503, Message: "error"})
	}

	osde.Identificador = idosde
	var resultUpdate, err, result = interfaces.UpdateOsde(osde)

	if result == true {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(resultUpdate)
}

// @Summary Eliminar una OSDE
// @Description Elimina la OSDE del sistema
// @Tags OSDE
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "OSDE"
// @Success 200 {object} models.OsdeResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/osde/{id} [delete]
func OsdeDelete(c *fiber.Ctx) error {
	idosde := c.Params("id")

	var resultDelete, err, result = interfaces.DeleteOsde(idosde)

	if result == true {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(resultDelete)
}
