package models

type CreateTenderRequest struct {

	// Полное название тендера
	Name string `json:"name"`

	// Описание тендера
	Description string `json:"description"`

	ServiceType TenderServiceType `json:"serviceType"`

	// Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId string `json:"organizationId"`

	// Уникальный slug пользователя.
	CreatorUsername string `json:"creatorUsername"`
}