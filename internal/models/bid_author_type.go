package models


import (
	"fmt"
)


// BidAuthorType : Тип автора
type BidAuthorType string

// List of BidAuthorType
const (
	ORGANIZATION BidAuthorType = "Organization"
	USER BidAuthorType = "User"
)

// AllowedBidAuthorTypeEnumValues is all the allowed values of BidAuthorType enum
var AllowedBidAuthorTypeEnumValues = []BidAuthorType{
	"Organization",
	"User",
}

// validBidAuthorTypeEnumValue provides a map of BidAuthorTypes for fast verification of use input
var validBidAuthorTypeEnumValues = map[BidAuthorType]struct{}{
	"Organization": {},
	"User": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BidAuthorType) IsValid() bool {
	_, ok := validBidAuthorTypeEnumValues[v]
	return ok
}

// NewBidAuthorTypeFromValue returns a pointer to a valid BidAuthorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBidAuthorTypeFromValue(v string) (BidAuthorType, error) {
	ev := BidAuthorType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for BidAuthorType: valid values are %v", v, AllowedBidAuthorTypeEnumValues)
}