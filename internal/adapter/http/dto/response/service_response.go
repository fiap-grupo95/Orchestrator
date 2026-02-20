package response

// ServiceResponse represents the response payload for service operations.
type ServiceResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
