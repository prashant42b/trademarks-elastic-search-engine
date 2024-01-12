package routeHandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/elastic-search-engine-task/database"
	"github.com/prashant42b/elastic-search-engine-task/internal/model"
	elasticsearchutils "github.com/prashant42b/elastic-search-engine-task/utils/elasticsearch_utils"
)

func HandleESSearch(c *fiber.Ctx) error {
	req := new(elasticsearchutils.Payload)
	err := c.BodyParser(&req)
	if err != nil {
		if err.Error() != "EOF" {
			return c.Status(400).SendString("Error parsing request body")
		}
	} else {

		results, err := elasticsearchutils.TrademarkSearch(req)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to perform Elasticsearch query: %v", err),
			})
		}

		// Respond with the JSON received from the Elasticsearch query
		return c.Status(200).SendString(results)

	}
	return c.Status(200).SendString("Worked")
}

func HandleSerialNumberQuerySearch(c *fiber.Ctx) error {

	db := database.DB

	var trademark model.Trademarks

	// Extract the search term from the url params
	searchTerm := c.Params("serialNumber")

	// Check if the search term is provided
	// if trademark.ID == uuid.Nil {
	// 	return c.Status(404).JSON(fiber.Map{"status": "Not found", "message": "Trademark not found", "data": nil})
	// }
	db.Find(&trademark, "serial_number = ?", searchTerm)

	// Return the search results
	return c.JSON(fiber.Map{"status": "success", "message": "Search results", "data": trademark})

}
