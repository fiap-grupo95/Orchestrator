package response

import "time"

// AdditionalRepairServiceResponse represents a service returned to the client.
type AdditionalRepairServiceResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

// AdditionalRepairPartsSupplyResponse represents a parts supply returned to the client.
type AdditionalRepairPartsSupplyResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity"`
}

// AdditionalRepairResponse encapsulates the data returned for a single additional repair.
type AdditionalRepairResponse struct {
	ID             string                                `json:"id"`
	ServiceOrderID string                                `json:"service_order_id"`
	Description    string                                `json:"description"`
	Status         string                                `json:"status"`
	Estimate       *EstimateResponse                     `json:"estimate"`
	CreatedAt      time.Time                             `json:"created_at"`
	UpdatedAt      time.Time                             `json:"updated_at"`
	Services       []AdditionalRepairServiceResponse     `json:"services"`
	PartsSupplies  []AdditionalRepairPartsSupplyResponse `json:"parts_supplies"`
}

// OperationMessageResponse represents a simple operation feedback payload.
type OperationMessageResponse struct {
	Message string `json:"message"`
}
