package models

// TenderStatus : Статус тендер
type TenderStatus string

// List of TenderStatus
const (
	TENDER_CREATED TenderStatus = "CREATED"
	TENDER_PUBLISHED TenderStatus = "PUBLISHED"
	TENDER_CLOSED TenderStatus = "CLOSED"
)