package request

type StartInput struct {
	CustomerID string   `json:"customer_id"`
	VehicleID  string   `json:"vehicle_id"`
	Items      []string `json:"items"`
}
