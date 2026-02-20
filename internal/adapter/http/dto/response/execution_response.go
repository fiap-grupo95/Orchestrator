package response

import "time"

type ExecutionResponse struct {
	ID             string     `json:"id"`
	ServiceOrderID string     `json:"service_order_id"`
	Status         string     `json:"status"`
	StartedAt      *time.Time `json:"started_at"`
	FinishedAt     *time.Time `json:"finished_at"`
}
