package models

import "fmt"

// TenderServiceType : Вид услуги, к которой относиться тендер
type TenderServiceType string

// List of TenderServiceType
const (
	CONSTRUCTION TenderServiceType = "Construction"
	DELIVERY TenderServiceType = "Delivery"
	MANUFACTURE TenderServiceType = "Manufacture"
)

// AllowedTenderServiceTypeEnumValues is all the allowed values of TenderServiceType enum
var AllowedTenderServiceTypeEnumValues = []TenderServiceType{
	"Construction",
	"Delivery",
	"Manufacture",
}

// validTenderServiceTypeEnumValue provides a map of TenderServiceTypes for fast verification of use input
var validTenderServiceTypeEnumValues = map[TenderServiceType]struct{}{
	"Construction": {},
	"Delivery": {},
	"Manufacture": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenderServiceType) IsValid() bool {
	_, ok := validTenderServiceTypeEnumValues[v]
	return ok
}

// NewTenderServiceTypeFromValue returns a pointer to a valid TenderServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenderServiceTypeFromValue(v string) (TenderServiceType, error) {
	ev := TenderServiceType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for TenderServiceType: valid values are %v", v, AllowedTenderServiceTypeEnumValues)
}