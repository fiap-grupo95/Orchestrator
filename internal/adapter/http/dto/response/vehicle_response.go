package response

// VehicleResponse represents the payload returned for vehicle operations.
type VehicleResponse struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Year       string `json:"year"`
	Plate      string `json:"plate"`
}

// VehicleCustomerSummary represents minimal customer information included with vehicle endpoints.
type VehicleCustomerSummary struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Document    string `json:"document"`
}

// VehiclesByCustomerResponse groups vehicles with their corresponding customer.
type VehiclesByCustomerResponse struct {
	Customer *VehicleCustomerSummary `json:"customer,omitempty"`
	Vehicles []VehicleResponse       `json:"vehicles"`
}
