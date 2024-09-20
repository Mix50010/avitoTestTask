package models

type EditTenderRequest struct {

	// Полное название тендера
	Name string `json:"name,omitempty"`

	// Описание тендера
	Description string `json:"description,omitempty"`

	ServiceType TenderServiceType `json:"serviceType,omitempty"`
}