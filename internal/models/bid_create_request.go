package models

type CreateBidRequest struct {

	// Полное название предложения
	Name string `json:"name"`

	// Описание предложения
	Description string `json:"description"`

	// Уникальный идентификатор тендера, присвоенный сервером.
	TenderId string `json:"tenderId"`

	AuthorType BidAuthorType `json:"authorType"`

	// Уникальный идентификатор автора предложения, присвоенный сервером.
	AuthorId string `json:"authorId"`
}