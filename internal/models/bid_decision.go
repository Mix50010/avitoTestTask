package models


import (
	"fmt"
)


// BidDecision : Решение по предложению
type BidDecision string

// List of BidDecision
const (
	APPROVED BidDecision = "Approved"
	REJECTED BidDecision = "Rejected"
)

// AllowedBidDecisionEnumValues is all the allowed values of BidDecision enum
var AllowedBidDecisionEnumValues = []BidDecision{
	"Approved",
	"Rejected",
}

// validBidDecisionEnumValue provides a map of BidDecisions for fast verification of use input
var validBidDecisionEnumValues = map[BidDecision]struct{}{
	"Approved": {},
	"Rejected": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BidDecision) IsValid() bool {
	_, ok := validBidDecisionEnumValues[v]
	return ok
}

// NewBidDecisionFromValue returns a pointer to a valid BidDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBidDecisionFromValue(v string) (BidDecision, error) {
	ev := BidDecision(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for BidDecision: valid values are %v", v, AllowedBidDecisionEnumValues)
}