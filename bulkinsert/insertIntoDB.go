package bulkinsert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/prashant42b/trademarks-elastic-search-engine/config"
	"github.com/prashant42b/trademarks-elastic-search-engine/database"
	"github.com/prashant42b/trademarks-elastic-search-engine/internal/model"
)

func BulkInsertJsonIntoDB() {

	var DB = database.DB
	var count = 0

	// Check if DB is nil
	if DB == nil {
		log.Fatal("Database connection is nil")
		return
	}

	// Read the JSON data
	byteValue, err := ioutil.ReadFile(config.JSON_FILE_PATH)
	if err != nil {
		log.Fatal(err)
		return
	}

	var trademarksList []model.Trademarks
	err = json.Unmarshal(byteValue, &trademarksList)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, data := range trademarksList {
		t, err := time.Parse("20060102", data.ApplicationDate)
		if err != nil {
			continue
		}
		x, err := time.Parse("20060102", data.RegistrationDate)
		if err != nil {
			continue
		}

		data.ApplicationDate = fmt.Sprintf("%d", t.Unix())
		data.RegistrationDate = fmt.Sprintf("%d", x.Unix())
		err = DB.Create(&data).Error
		if err != nil {
			log.Printf("Error inserting data for patent %s: %v", data.SerialNumber, err)
		} else {
			count += 1
			fmt.Printf("Data inserted successfully for trademark %s!\n", data.SerialNumber)
		}
	}
	log.Println("total records inserted: ", count)

}
