package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/prashant42b/elastic-search-engine-task/config"
	"github.com/prashant42b/elastic-search-engine-task/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	// p := config.Config("DB_PORT")
	p := config.PORT
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error parsing PORT NO")
		return
	}

	// dataSource connection URL to connect to Postgres Database
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.HOST, port, config.USER, config.PASSWORD, config.NAME)

	// Connect to the DB and initialize the DB variable
	gormConfig := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	}

	DB, err = gorm.Open(postgres.Open(dataSource), gormConfig)

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

}

func AutoMigrateDB() {

	//migrating gorm model to db
	err := DB.AutoMigrate(&model.Trademarks{})
	if err != nil {
		fmt.Printf("Error migrating DB: %v\n", err)
		return
	}
	fmt.Println("Database Migrated")
}
