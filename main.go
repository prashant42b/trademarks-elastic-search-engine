package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/elastic-search-engine-task/bulkinsert"
	"github.com/prashant42b/elastic-search-engine-task/config"
	"github.com/prashant42b/elastic-search-engine-task/database"
	"github.com/prashant42b/elastic-search-engine-task/utils"
	//"github.com/prashant42b/elastic-search-engine-task/utils/conversion_utils"
)

func main() {
	//start a new fiber app
	utils.ImportENV()
	config.LoadConfig()

	app := fiber.New()

	//Unzips XML file from archive
	//conversion_utils.UnzipXML(config.ZIP_NAME)

	// Extract xml to json
	//conversion_utils.CleanAndConvert(config.XML_PATH)

	//Migrates gorm model to postgres db
	database.ConnectDB()
	//database.AutoMigrateDB()

	//Bulk inserts jsonData into postgres table
	bulkinsert.BulkInsertJsonIntoDB()

	//Redirects the app to endpoints defined for the api
	//router.SetupRoutes(app)

	//Listens on port 3000
	app.Listen(":3000")

}
