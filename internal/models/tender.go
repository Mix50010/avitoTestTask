package models

import "time"

// Tender - Информация о тендере
type Tender struct {

	// Уникальный идентификатор тендера, присвоенный сервером.
	Id string `json:"id"`

	// Полное название тендера
	Name string `json:"name"`

	// Описание тендера
	Description string `json:"description"`

	ServiceType TenderServiceType `json:"serviceType"`

	Status TenderStatus `json:"status"`

	// Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId string `json:"organizationId"`

	// Номер версии посел правок
	Version int32 `json:"version"`

	// Серверная дата и время в момент, когда пользователь отправил тендер на создание. Передается в формате RFC3339. 
	CreatedAt time.Time `json:"createdAt"`
}