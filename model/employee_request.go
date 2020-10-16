package model

// EmployeeCreatePayload...
type EmployeeCreatePayload struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
}

// EmployeeUpdatePayload...
type EmployeeUpdatePayload struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
}
