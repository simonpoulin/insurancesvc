package model

// InsuranceTypeResponse...
type InsuranceTypeResponse struct {
	Name      string          `json:"name"`
	CarSizeID CarSizeResponse `json:"car_size"`
}
