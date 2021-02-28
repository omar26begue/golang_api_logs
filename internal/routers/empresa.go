package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-rest-api-logs/internal/handlers"
)

func EmpresaRouterV1(app fiber.Router) {
	routerEmpresa := app.Group("/empresas")
	routerEmpresa.Get("", handlers.EmpresasGetAll)
	routerEmpresa.Post("", handlers.EmpresasCreate)
	routerEmpresa.Patch("/:id", handlers.EmpresaUpdate)
	routerEmpresa.Delete("/:id", handlers.EmpresaDelete)
}
