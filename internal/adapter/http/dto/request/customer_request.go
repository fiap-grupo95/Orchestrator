package request

// CustomerCreateRequest represents the payload to create a new customer.
type CustomerCreateRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Document    string `json:"document" binding:"required"`
}

// CustomerUpdateRequest represents the payload to update an existing customer.
// All fields are optional to support partial updates.
type CustomerUpdateRequest struct {
	FullName    string `json:"full_name" binding:"omitempty"`
	Email       string `json:"email" binding:"omitempty,email"`
	PhoneNumber string `json:"phone_number" binding:"omitempty"`
}
