package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-api-logs/internal/interfaces"
)

// @Summary Devuelve todos los códigos de respuestas
// @Description Devuelve todos los códigos de respuesta de las trazas.
// @Tags Códigos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.EmpresasResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Security ApiKeyAuth
// @Router /v1/codes [get]
func CodesResponseGetAll(c *fiber.Ctx) error {
	var codes = interfaces.GetAllCodesResponse()

	return c.Status(fiber.StatusOK).JSON(codes)
}
