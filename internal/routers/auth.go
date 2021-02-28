package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-api-logs/internal/handlers"
)

//AuthRouter Autentificación mediante usuario
func AuthRouter(app fiber.Router) {
	app.Post("/login", handlers.LoginUsers)
}
