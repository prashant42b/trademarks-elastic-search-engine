package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/elastic-search-engine-task/config"
	"github.com/prashant42b/elastic-search-engine-task/utils"
)

func main() {
	//start a new fiber app
	utils.ImportENV()
	config.LoadConfig()

	app := fiber.New()

	//Unzips XML file from archive
	utils.UnzipXML()

	//database.ConnectDB()
	//router.SetupRoutes(app)

	//port no 3000
	app.Listen(":3000")
}
