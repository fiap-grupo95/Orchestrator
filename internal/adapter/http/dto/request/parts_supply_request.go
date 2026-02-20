package request

// PartsSupplyCreateRequest represents payload for creating a parts supply record.
type PartsSupplyCreateRequest struct {
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
	QuantityTotal   int     `json:"quantity_total" binding:"required"`
	QuantityReserve int     `json:"quantity_reserve" binding:"omitempty"`
}

// PartsSupplyUpdateRequest represents payload for updating a parts supply record.
type PartsSupplyUpdateRequest struct {
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
	QuantityTotal   int     `json:"quantity_total" binding:"required"`
	QuantityReserve int     `json:"quantity_reserve" binding:"required"`
}
