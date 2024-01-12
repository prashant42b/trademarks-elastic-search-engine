# USPTO Trademarks Search Engine

This project is a Go web application developed using the Fiber framework. It serves as a platform for searching and retrieving trademark information sourced from USPTO bulk data based on various criteria like Mark Identifcation, Attorney Names, Owners, Serial number, Class Codes, and Application Date.


## Table of Contents

- [Overview](#overview)
- [Usage](#usage)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Search Architecture](#search-architecture)
  - [Flow Diagram](#flow-diagram)
- [Features](#features)
- [Data Source](#data-source)
- [Extracting Data](#extracting-data)
- [Bulk Insertion](#bulk-insertion)
- [Search](#search)
- [Postman Documentation](#postman-documentation)

## Overview

This is a data-driven application that leverages information obtained from the United States Patent and Trademark Office (USPTO) dataset. It utilizes USPTO daily trademark files to provide valuable insights, search capabilities, and retrieval of trademark information.

## Usage

### Prerequisites

The following prerequisites are required to run this application:

- GoLang
- ElasticSearch
- PostgreSQL


### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/prashant42b/elastic-search-engine-task.git
   cd elastic-search-engine-task

   ```
2. Run 
      ```
      go mod download
      go run main.go
      ```
      
## Search Architecture
### Flow Diagram
<img width="705" alt="image" src="https://github.com/prashant42b/elastic-search-engine-task/assets/63443918/0bd994b0-04ec-47ba-84aa-9a6c3afa784e">

## Features

- Search trademarks by mark-identification, serial number, attorney name(s), owner(s), application date, and class code(s).
- Efficiently parse and store USPTO trademark data from XML to JSON.
- The project integrates with a PostgreSQL database using GORM (ORM) to store and retrieve trademark information.

## Data Source
The USPTO data is sourced from the official United States Patent and Trademark Office database.

## Extracting Data
- Files: xmlFromArchive.go and xmlToJSON.go
- The extraction process entails a script that decompresses the daily data file, extracting the XML file into a designated folder. 
- The next step involves utilizing the encoding/xml and encoding/json packages to parse all the extracted fields into JSON.

```
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
```

```
type Owner struct {
	Name string `xml:"party-name"`
}
```

## Bulk Insertion
- Files: insertIntoDB.go and insertIntoESDB.go
- Two utils have been created to facilitate Bulk insertion of json data from the converted_data.json file.
- Bulk insertion into PostgreSQL DB: insertIntoDB.go makes use of a GORM model to handle this functionality.
- Bulk insertion into Elastic Search: insertIntoESDB.go employs a strategy to bulk insert jsonData into the Elastic Search Index (trademarks)

## Search
The search engine uses Elastic Search and enables users to search and retrieve trademark information sourced from USPTO bulk data based on various criteria like Mark Identifcation, Attorney Names, Owners, Serial number, Class Codes, and Application Date.


## Postman Documentation 
- [Postman Documentation](https://documenter.getpostman.com/view/30488190/2s9YsNcpXA) 
Please refer to the above mentioned API documentation.
