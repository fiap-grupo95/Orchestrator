package request

// VehicleCreateRequest represents the payload to create a new vehicle.
type VehicleCreateRequest struct {
	CustomerID string `json:"customer_id" binding:"required"`
	Brand      string `json:"brand" binding:"required"`
	Model      string `json:"model" binding:"required"`
	Year       string `json:"year" binding:"required"`
	Plate      string `json:"plate" binding:"required"`
}

// VehicleUpdateRequest represents the payload to partially update a vehicle.
type VehicleUpdateRequest struct {
	CustomerID *string `json:"customer_id,omitempty"`
	Brand      *string `json:"brand,omitempty"`
	Model      *string `json:"model,omitempty"`
	Year       *string `json:"year,omitempty"`
	Plate      *string `json:"plate,omitempty"`
}
