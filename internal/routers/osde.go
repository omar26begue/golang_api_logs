package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-api-logs/internal/handlers"
)

func OsdeRouterV1(app fiber.Router) {
	routerOsde := app.Group("/osde")
	routerOsde.Get("", handlers.OsdeGetAll)
	routerOsde.Get("/:id", handlers.OsdeGetFindOne)
	routerOsde.Post("", handlers.OsdeCreate)
	routerOsde.Patch("/:id", handlers.OsdeUpdate)
	routerOsde.Delete("/:id", handlers.OsdeDelete)
}

func OsdeRouterV2(app fiber.Router) {
	routerOsde := app.Group("/osde")
	routerOsde.Get("/", handlers.OsdeGetAllv2)
}
