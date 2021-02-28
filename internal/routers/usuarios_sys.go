package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-api-logs/internal/handlers"
)

func CreateUsers(app fiber.Router) {
	app.Post("users", handlers.CreateUsers)
}
