package trademarkRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/elastic-search-engine-task/internal/routeHandler"
)

func SetupTrademarkRoutes(router fiber.Router) {

	query := router.Group("/query")
	query.Get("/:serialNumber", routeHandler.HandleSerialNumberQuerySearch)

	search := router.Group("/search")
	search.Post("/", routeHandler.HandleESSearch)

}
