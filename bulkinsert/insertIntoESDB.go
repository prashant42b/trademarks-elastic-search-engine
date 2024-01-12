package bulkinsert

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/prashant42b/elastic-search-engine-task/config"
	"github.com/prashant42b/elastic-search-engine-task/database"
)

func BulkInsertJsonIntoESDB() {

	var es = database.GetElasticClient()

	filePath := config.JSON_FILE_PATH

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %s", err)
	}

	// Unmarshal JSON data
	var jsonData []map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}

	// Bulk index request
	var bulkRequestBody bytes.Buffer

	for _, document := range jsonData {

		sanitizeJSONToLowercase(document)

		// Exclude StatusCode, MarkDrawingCode, FilingDate, RegistrationDate and RegistrationNumber fields
		delete(document, "status_code")
		delete(document, "mark_drawing_code")
		delete(document, "filing_date")
		delete(document, "registration_date")
		delete(document, "registration_number")

		metaData := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": config.ES_INDEX_NAME,
			},
		}

		if err := json.NewEncoder(&bulkRequestBody).Encode(metaData); err != nil {
			log.Fatalf("Error encoding metadata: %s", err)
		}

		if err := json.NewEncoder(&bulkRequestBody).Encode(document); err != nil {
			log.Fatalf("Error encoding document: %s", err)
		}
	}

	res, err := es.Bulk(bytes.NewReader(bulkRequestBody.Bytes()), es.Bulk.WithContext(context.Background()))
	if err != nil {
		log.Fatalf("Error performing bulk insert: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var buf bytes.Buffer
		io.Copy(&buf, res.Body)
		log.Fatalf("Error response: %s", buf.String())
	}

	log.Println("Bulk insertion completed successfully.")
}

func sanitizeJSONToLowercase(document map[string]interface{}) {
	for key, value := range document {
		switch v := value.(type) {
		case string:
			document[key] = strings.ToLower(v)
		case map[string]interface{}:
			sanitizeJSONToLowercase(v)
		case []interface{}:
			for _, item := range v {
				if nested, ok := item.(map[string]interface{}); ok {
					sanitizeJSONToLowercase(nested)
				}
			}
		}
	}
}
