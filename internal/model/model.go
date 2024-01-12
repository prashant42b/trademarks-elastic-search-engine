package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Trademarks struct {
	gorm.Model
	SerialNumber       string         `gorm:"primaryKey" json:"serial_number"`
	FilingDate         string         `json:"filing_date"`
	StatusCode         string         `json:"status_code"`
	MarkIdentification string         `json:"mark_identification"`
	MarkDrawingCode    string         `json:"mark_drawingcode"`
	AttorneyNames      string         `json:"attorney_names"`
	Owners             pq.StringArray `json:"owners" gorm:"type:text[]"`
	ApplicationDate    string         `json:"application_date"`
	RegistrationNumber string         `json:"registration_number"`
	ClassCode          string         `json:"class_code"`
	RegistrationDate   string         `json:"registration_date"`
}
