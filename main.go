package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	_ "go-rest-api-logs/docs"
	"go-rest-api-logs/internal/database"
	"go-rest-api-logs/internal/middlewares"
	"go-rest-api-logs/internal/parse"
	"go-rest-api-logs/internal/routers"
)

// @title API Logs - GET
// @version 1.0.0
// @description Prueba API
// @termsOfService http://swagger.io/terms/

// @contact.name Omar Isalgue Begue
// @contact.email omar26begue@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// leyendo las variables de entorno
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// base de datos
	database.GetConnection()

	// programacion de tareas
	c := cron.New()
	c.Start()
	//c.AddFunc("*/10 * * * * *", parse.SquidParseCreateFolder)
	defer c.Stop()

	parse.SquidParseCreateFolder()

	app := fiber.New()
	//app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.Handler)

	// monitor
	app.Get("/monitor", monitor.New())

	// versiones
	apiRouter := app.Group("/api", logger.New())
	{
		authRouter := apiRouter.Group("/auth")
		{
			routers.AuthRouter(authRouter)
		}

		// administracion
		adminRouter := apiRouter.Group("/", middlewares.Protected())
		{
			routers.CreateUsers(adminRouter)
		}

		// version 1
		v1 := apiRouter.Group("/v1", middlewares.Protected())
		{
			routers.OsdeRouterV1(v1)
			routers.EmpresaRouterV1(v1)
			routers.CodesRouterV1(v1)
		}

		// version 2
		v2 := apiRouter.Group("/v2", middlewares.Protected())
		{
			routers.OsdeRouterV2(v2)
		}
	}

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":" + viper.GetString("PORT"))
}
