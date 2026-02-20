package entities

import (
	"github.com/daniloAleite/orchestrator/internal/domain/valueobject"
	"time"
)

type ServiceOrder struct {
	ID                string                         `json:"id"`
	CustomerID        string                         `json:"customer_id"`
	Customer          *Customer                      `json:"customer"`
	VehicleID         string                         `json:"vehicle_id"`
	Vehicle           *Vehicle                       `json:"vehicle"`
	Status            valueobject.ServiceOrderStatus `json:"service_order_status"`
	Estimate          *Estimate                      `json:"estimate,omitempty"`
	Execution         *Execution                     `json:"execution,omitempty"`
	CreatedAt         *time.Time                     `json:"created_at,omitempty"`
	UpdatedAt         *time.Time                     `json:"updated_at,omitempty"`
	AdditionalRepairs []AdditionalRepair             `json:"additional_repairs,omitempty"`
	PartsSupplies     []PartsSupply                  `json:"parts_supplies,omitempty"`
	Services          []Service                      `json:"services,omitempty"`
}

func (s *ServiceOrder) IsDiagnosisPending() bool {
	return len(s.Services) == 0 && len(s.PartsSupplies) == 0
}
