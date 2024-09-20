package models

import "time"

// Bid - Информация о предложении
type Bid struct {

	// Уникальный идентификатор предложения, присвоенный сервером.
	Id string `json:"id"`

	// Полное название предложения
	Name string `json:"name"`

	// Описание предложения
	Description string `json:"description"`

	Status BidStatus `json:"status"`

	// Уникальный идентификатор тендера, присвоенный сервером.
	TenderId string `json:"tenderId"`

	AuthorType BidAuthorType `json:"authorType"`

	// Уникальный идентификатор автора предложения, присвоенный сервером.
	AuthorId string `json:"authorId"`

	// Номер версии посел правок
	Version int32 `json:"version"`

	// Серверная дата и время в момент, когда пользователь отправил предложение на создание. Передается в формате RFC3339. 
	CreatedAt time.Time `json:"createdAt"`
}