package entities

type Estimate struct {
	ID                 string  `json:"id"`
	ServiceOrderID     string  `json:"service_order_id"`
	AdditionalRepairID string  `json:"additional_repair_id"`
	Value              float64 `json:"value"`
	Status             string  `json:"status"`
}
