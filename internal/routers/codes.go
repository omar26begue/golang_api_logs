package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-api-logs/internal/handlers"
)

func CodesRouterV1(app fiber.Router) {
	routerCodes := app.Group("/codes")
	routerCodes.Get("", handlers.CodesResponseGetAll)
}
