package router

import (
	"github.com/gofiber/fiber/v2"

	trademarkRoutes "github.com/prashant42b/elastic-search-engine-task/internal/routes"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	//Setup node routers
	trademarkRoutes.SetupTrademarkRoutes(api)

}
