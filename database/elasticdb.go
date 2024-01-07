package database

import (
	"log"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

var esClient *elasticsearch.Client
var esClientOnce sync.Once
var ES_INDEX_NAME string

func initializeElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{viper.GetString("ES_HOST")},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	ES_INDEX_NAME = viper.GetString("ES_INDEX_NAME")

	esClient = client
}

func InitElasticsearch() {
	esClientOnce.Do(initializeElasticsearch)
}

func GetElasticClient() *elasticsearch.Client {
	return esClient
}
