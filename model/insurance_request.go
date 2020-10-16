package model

// InsuranceCreatePayload...
type InsuranceCreatePayload struct {
	Price       int                  `json:"price"`
	Value       int                  `json:"value"`
	Description string               `json:"desc"`
	Type        InsuranceTypePayload `json:"type"`
}

// InsuranceUpdatePayload...
type InsuranceUpdatePayload struct {
	Price       int                  `json:"price"`
	Value       int                  `json:"value"`
	Description string               `json:"desc"`
	Type        InsuranceTypePayload `json:"type"`
}
