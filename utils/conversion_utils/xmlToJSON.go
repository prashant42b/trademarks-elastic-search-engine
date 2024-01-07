package conversion_utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/prashant42b/elastic-search-engine-task/config"
)

// type Attorney struct {
// 	Name string `xml:"attorney-name"`
// }

type Owner struct {
	Name string `xml:"party-name"`
}

type CaseFile struct {
	SerialNumber       string  `xml:"serial-number"`
	FilingDate         string  `xml:"case-file-header>filing-date"`
	StatusCode         string  `xml:"case-file-header>status-code"`
	MarkIdentification string  `xml:"case-file-header>mark-identification"`
	MarkDrawingCode    string  `xml:"case-file-header>mark-drawing-code"`
	AttorneyNames      string  `xml:"case-file-header>attorney-name"`
	Owners             []Owner `xml:"case-file-owners>case-file-owner"`
	ApplicationDate    string  `xml:"transaction-date"`
	RegistrationNumber string  `xml:"registration-number"`
	ClassCode          string  `xml:"classifications>classification>international-code"`
	RegistrationDate   string  `xml:"case-file-header>registration-date"`
}

func sanitize(elem string) string {
	if elem != "" {
		return elem
	}
	return ""
}

func CleanAndConvert(xmlFilePath string) {
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Printf("Error opening XML file: %v\n", err)
		return
	}
	defer xmlFile.Close()

	var root struct {
		CaseFiles []CaseFile `xml:"application-information>file-segments>action-keys>case-file"`
	}

	decoder := xml.NewDecoder(xmlFile)
	err = decoder.Decode(&root)
	if err != nil {
		fmt.Printf("Error decoding XML: %v\n", err)
		return
	}

	var jsonData []map[string]interface{}
	for _, caseFile := range root.CaseFiles {
		caseData := map[string]interface{}{
			"serial_number":       sanitize(caseFile.SerialNumber),
			"filing_date":         sanitize(caseFile.FilingDate),
			"status_code":         sanitize(caseFile.StatusCode),
			"mark_identification": sanitize(caseFile.MarkIdentification),
			"mark_drawingcode":    sanitize(caseFile.MarkDrawingCode),
			"attorney_names":      sanitize(caseFile.AttorneyNames),
			"owners":              []string{},
			"application_date":    sanitize(caseFile.ApplicationDate),
			"registration_number": sanitize(caseFile.RegistrationNumber),
			"class_code":          sanitize(caseFile.ClassCode),
			"registration_date":   sanitize(caseFile.RegistrationDate),
		}

		// for _, attorney := range caseFile.AttorneyNames {
		// 	caseData["attorney_names"] = append(caseData["attorney_names"].([]string), sanitize(attorney.Name))
		// }

		for _, owner := range caseFile.Owners {
			caseData["owners"] = append(caseData["owners"].([]string), sanitize(owner.Name))
		}

		jsonData = append(jsonData, caseData)
	}

	jsonFile, err := os.Create(config.JSON_FILE_PATH)
	if err != nil {
		fmt.Printf("Error creating JSON file: %v\n", err)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(jsonData)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}
}
