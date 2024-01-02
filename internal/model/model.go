package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Trademark struct {
	gorm.Model
	SerialNumber    uuid.UUID `gorm:"type:uuid"`
	TrademarkName   string    `json:"TrademarkName"`
	AttorneyName    string    `json:"AttorneyName"`
	Owner           string    `json:"Owner"`
	LawFirm         string    `json:"LawFirm"`
	ApplicationDate string    `json:"ApplicationDate"`
	ClassCode       string    `json:"ClassCode"`
}
