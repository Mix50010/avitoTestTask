package models

import (
	"fmt"
)


// BidStatus : Статус предложения
type BidStatus string

// List of BidStatus
const (
	BID_CREATED BidStatus = "Created"
	BID_PUBLISHED BidStatus = "Published"
	BID_CANCELED BidStatus = "Canceled"
)

// AllowedBidStatusEnumValues is all the allowed values of BidStatus enum
var AllowedBidStatusEnumValues = []BidStatus{
	"Created",
	"Published",
	"Canceled",
}

// validBidStatusEnumValue provides a map of BidStatuss for fast verification of use input
var validBidStatusEnumValues = map[BidStatus]struct{}{
	"Created": {},
	"Published": {},
	"Canceled": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BidStatus) IsValid() bool {
	_, ok := validBidStatusEnumValues[v]
	return ok
}

// NewBidStatusFromValue returns a pointer to a valid BidStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBidStatusFromValue(v string) (BidStatus, error) {
	ev := BidStatus(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for BidStatus: valid values are %v", v, AllowedBidStatusEnumValues)
}