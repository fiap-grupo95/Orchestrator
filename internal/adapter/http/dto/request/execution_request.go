package request

type ExecutionRequest struct {
	ID             string   `json:"id"`
	ServiceOrderID string   `json:"service_order_id"`
}