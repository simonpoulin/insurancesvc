package model

// CarSizeCreatePayload...
type CarSizeCreatePayload struct {
	Description string `json:"desc"`
	MinSeat     int    `json:"min_seat"`
	MaxSeat     int    `json:"max_seat"`
}

// CarSizeUpdatePayload...
type CarSizeUpdatePayload struct {
	Description string `json:"desc"`
	MinSeat     int    `json:"min_seat"`
	MaxSeat     int    `json:"max_seat"`
}
