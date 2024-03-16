package elasticsearchutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/prashant42b/trademarks-elastic-search-engine/database"
)

type Payload struct {
	SearchTerm      string   `json:"search_term"`
	ApplicationDate string   `json:"application_date"`
	ClassCodes      []string `json:"class_codes"`
	Owners          []string `json:"owners"`
	AttorneyNames   string   `json:"attorney_names"`
}

type DesiredQuery struct {
	Query QueryStruct `json:"query"`
	Aggs  AggsStruct  `json:"aggs"`
}

type QueryStruct struct {
	Bool BoolStruct `json:"bool"`
}

type BoolStruct struct {
	Should []map[string]interface{} `json:"should"`
	Filter []map[string]interface{} `json:"filter"`
}

type AggsStruct struct {
	ApplicationDateTerms map[string]interface{} `json:"application_date_terms"`
	ClassCodes           map[string]interface{} `json:"class_codes"`
	Owners               map[string]interface{} `json:"owners"`
	AttorneyNames        map[string]interface{} `json:"attorney_names"`
}

func TrademarkSearch(req *Payload) (string, error) {
	searchTerm := req.SearchTerm
	applicationDate := req.ApplicationDate
	classCodes := req.ClassCodes
	owners := req.Owners
	attorneyNames := req.AttorneyNames

	var es = database.GetElasticClient()

	aggregations := map[string]interface{}{
		"application_date_terms": map[string]interface{}{
			"terms": map[string]interface{}{
				"field": "application_date.keyword",
			},
		},
		"class_codes": map[string]interface{}{
			"terms": map[string]interface{}{
				"field": "class_code.keyword",
			},
		},
		"owners": map[string]interface{}{
			"terms": map[string]interface{}{
				"field": "owners.keyword",
			},
		},
		"attorney_names": map[string]interface{}{
			"terms": map[string]interface{}{
				"field": "attorney_names.keyword",
			},
		},
	}

	// Default query structure
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"mark_identification": searchTerm,
						},
					},
				},
				"filter": []map[string]interface{}{},
			},
		},
		"aggs": aggregations,
	}

	// Handle serial_number condition
	if len(searchTerm) == 8 && isNumeric(searchTerm) {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = []map[string]interface{}{
			{
				"match": map[string]interface{}{
					"serial_number": searchTerm,
				},
			},
		}
	}

	// Filters
	if applicationDate != "" {
		addTermFilter(query, "application_date.keyword", applicationDate)
	}

	if len(classCodes) > 0 {
		addTermsFilter(query, "class_code.keyword", classCodes)
	}

	if len(owners) > 0 {
		addTermsFilter(query, "owners.keyword", owners)
	}

	if attorneyNames != "" {
		addTermFilter(query, "attorney_names.keyword", attorneyNames)
	}

	// Convert the query to DesiredQuery structure
	desiredQuery := DesiredQuery{
		Query: QueryStruct{
			Bool: BoolStruct{
				Should: query["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"].([]map[string]interface{}),
				Filter: query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"].([]map[string]interface{}),
			},
		},
		Aggs: AggsStruct{
			ApplicationDateTerms: aggregations["application_date_terms"].(map[string]interface{}),
			ClassCodes:           aggregations["class_codes"].(map[string]interface{}),
			Owners:               aggregations["owners"].(map[string]interface{}),
			AttorneyNames:        aggregations["attorney_names"].(map[string]interface{}),
		},
	}

	// Convert the desiredQuery to JSON
	jsonQuery, err := json.Marshal(desiredQuery)
	if err != nil {
		return "", fmt.Errorf("error encoding the search query: %v", err)
	}

	fmt.Println(string(jsonQuery)) // Optional: Print the generated JSON

	// ES Search
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return "", fmt.Errorf("error encoding the search query: %v", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("trademarks"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return "", fmt.Errorf("error performing the search request: %v", err)
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	if res.IsError() {
		return "", fmt.Errorf("elasticsearch request failed with status code: %d", res.StatusCode)
	}

	return string(responseBody), nil
}

// Helper function to check if a string is numeric
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Helper function to add term filter to the query
func addTermFilter(query map[string]interface{}, field string, value string) {
	filters := query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"].([]map[string]interface{})
	filters = append(filters, map[string]interface{}{
		"term": map[string]interface{}{
			field: value,
		},
	})
	query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = filters
}

// Helper function to add terms filter to the query
func addTermsFilter(query map[string]interface{}, field string, values []string) {
	filters := query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"].([]map[string]interface{})
	filters = append(filters, map[string]interface{}{
		"terms": map[string]interface{}{
			field: values,
		},
	})
	query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = filters
}
