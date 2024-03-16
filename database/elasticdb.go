package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/prashant42b/trademarks-elastic-search-engine/config"
)

var esClient *elasticsearch.Client
var esClientOnce sync.Once
var ES_INDEX_NAME string

func initializeElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{config.ES_HOST},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}
	fmt.Println("Connection Opened to ES Client")

	esClient = client
}

func EstablishESConnection() {
	esClientOnce.Do(initializeElasticsearch)
}

func GetElasticClient() *elasticsearch.Client {
	return esClient
}
