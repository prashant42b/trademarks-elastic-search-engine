package trademarkRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/trademarks-elastic-search-engine/internal/routeHandler"
)

func SetupTrademarkRoutes(router fiber.Router) {

	query := router.Group("/query")
	query.Get("/:serialNumber", routeHandler.HandleSerialNumberQuerySearch)

	search := router.Group("/search")
	search.Post("/", routeHandler.HandleESSearch)

}
