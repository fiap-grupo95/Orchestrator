package response

// PartsSupplyResponse represents a parts supply payload returned to clients.
type PartsSupplyResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	QuantityTotal   int     `json:"quantity_total"`
	QuantityReserve int     `json:"quantity_reserve"`
}
