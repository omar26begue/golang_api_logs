package handlers

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go-rest-api-logs/internal/interfaces"
	"go-rest-api-logs/internal/models"
	"time"
)

// @Summary Inicia el registro de sessión de un usuario
// @Description Devuelve un token válido si la solicitud del usuario es correcta.
// @Tags Autentificación
// @ID login-users
// @Accept  json
// @Produce  json
// @Param param body models.UsuariosSysRequest true "Registro"
// @Success 200 {object} models.UsuariosSysResponse
// @Failure 400 {object} models.HTTPError400
// @Failure 401 {object} models.HTTPError401
// @Failure 500 {object} models.HTTPError500
// @Router /auth/login [post]
func LoginUsers(c *fiber.Ctx) error {
	login := new(models.UsuariosSys)
	if err := c.BodyParser(&login); err != nil {
		return c.Status(503).JSON(models.HTTPError400{Code: 503, Message: "error"})
	}

	var resultUsers, count = interfaces.GetFindOneUsers(login.Login)

	if count == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&models.HTTPError400{Code: 404, Message: "Su usuario no existe en el sistema."})
	}

	// comparando la contraseña
	if login.Password != resultUsers.Password {
		return c.Status(fiber.StatusBadRequest).JSON(&models.HTTPError400{Code: 400, Message: "Contraseña incorrecta."})
	}

	/*if password != resultUsers.Password {
		return c.Status(fiber.StatusBadRequest).JSON(&models.HTTPError400{Code: 400, Message: "Contraseña incorrecta."})
	}*/

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["identificador"] = resultUsers.Identificador
	claims["login"] = resultUsers.Login
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT")))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.HTTPError400{Code: 400, Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"idetificador": resultUsers.Identificador, "login": resultUsers.Login, "is_admin": resultUsers.IsAdmin, "token": tokenString})
}
