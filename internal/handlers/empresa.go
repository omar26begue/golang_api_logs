package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-rest-api-logs/internal/interfaces"
	"go-rest-api-logs/internal/models"
)

// @Summary Devuelve todas las empresas
// @Description Devuelve un listado de todas las empresas en el sistemas.
// @Tags Empresas
// @Accept  json
// @Produce  json
// @Success 200 {array} models.EmpresasResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/empresas [get]
func EmpresasGetAll(c *fiber.Ctx) error {
	var empresas = interfaces.GetAllEmpresas()

	return c.Status(fiber.StatusOK).JSON(empresas)
}

// @Summary Crea una empresa
// @Description Crea una empresa con los datos solicitados
// @Tags Empresas
// @Accept  json
// @Produce  json
// @Param empresa body models.EmpresasRequest true "Adicionar empresa"
// @Success 201 {object} models.EmpresasResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/empresas [post]
func EmpresasCreate(c *fiber.Ctx) error {
	empresas := new(models.EmpresasRequest)
	fmt.Println(c.FormValue("nombre_empresa"))

	if err := c.BodyParser(&empresas); err != nil {
		return c.Status(503).JSON(models.HTTPError400{Code: 503, Message: "error"})
	}

	var tempEmpresa = models.Empresas{Identificador: uuid.New().String(), NombreEmpresa: empresas.NombreEmpresa, IdOsde: empresas.IdOsde, Activo: true}

	var resultEmpresa, err, result = interfaces.CreateEmpresa(&tempEmpresa)

	if result == true {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON(fiber.Map{"identificador": resultEmpresa.Identificador, "nombre_empresa": resultEmpresa.NombreEmpresa, "id_osde": resultEmpresa.IdOsde})
}

// @Summary Modifica una empresa
// @Description Actualiza la información de una empresa
// @Tags Empresas
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "Empresa"
// @Param empresa body models.EmpresasRequest true "Actualizar datos empresa"
// @Success 200 {object} models.EmpresasResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/empresas/{id} [patch]
func EmpresaUpdate(c *fiber.Ctx) error {
	idempresa := c.Params("id")

	if _, errUUID := uuid.Parse(idempresa); errUUID != nil {
		return c.Status(400).JSON(models.HTTPError400{Code: 400, Message: "Identificador no válido"})
	}

	empresa := new(models.EmpresasRequest)
	if err := c.BodyParser(&empresa); err != nil {
		return c.Status(503).JSON(models.HTTPError400{Code: 503, Message: "error"})
	}

	var resultUpdate, err, result = interfaces.UpdateEmpresa(*empresa, idempresa)

	if result == true {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(resultUpdate)
}

// @Summary Eliminar una empresa
// @Description Elimina una empresaactiva del sistema
// @Tags Empresas
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "Empresa"
// @Success 200 {object} models.OsdeResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/empresas/{id} [delete]
func EmpresaDelete(c *fiber.Ctx) error {
	idempresa := c.Params("id")

	var resultDelete, err, result = interfaces.DeleteEmpresa(idempresa)

	if result == true {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(resultDelete)
}
