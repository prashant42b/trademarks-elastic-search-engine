package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/elastic-search-engine-task/bulkinsert"
	"github.com/prashant42b/elastic-search-engine-task/config"
	"github.com/prashant42b/elastic-search-engine-task/database"
	"github.com/prashant42b/elastic-search-engine-task/router"
	"github.com/prashant42b/elastic-search-engine-task/utils"
	"github.com/prashant42b/elastic-search-engine-task/utils/conversion_utils"
	"github.com/spf13/viper"
)

func main() {

	//loads env data into config
	utils.ImportENV()
	config.LoadConfig()

	//start a new fiber app
	app := fiber.New()

	// Use the log package in Fiber middleware
	app.Use(func(c *fiber.Ctx) error {
		// Log the request method and path
		log.Printf("Request received - Method: %s, Path: %s", c.Method(), c.Path())

		// Continue to the next middleware or route handler
		return c.Next()
	})

	//Establishes connection to postgres db and elastic_search
	database.ConnectDB()
	database.EstablishESConnection()

	if viper.GetBool("XML_TO_JSON") {

		//Unzips XML file from archive
		conversion_utils.UnzipXML(config.ZIP_NAME, config.XML_TARGET_FOLDER)

		// Extract xml to json
		conversion_utils.CleanAndConvert(config.XML_PATH)

	}

	if viper.GetBool("AUTO_MIGRATE") {

		//Migrates gorm model to postgres db
		database.AutoMigrateDB()
	}

	if viper.GetBool("BULK_INSERT") {

		//Bulk inserts jsonData into postgres table
		bulkinsert.BulkInsertJsonIntoDB()
		//bulkinsert.BulkInsertJsonIntoESDB()
	}

	//Redirects the app to endpoints defined for the api
	router.SetupRoutes(app)

	//Listens on port 3000
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
