package models

type EditBidRequest struct {

	// Полное название предложения
	Name string `json:"name,omitempty"`

	// Описание предложения
	Description string `json:"description,omitempty"`
}
