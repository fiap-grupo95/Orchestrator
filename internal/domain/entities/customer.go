package entities

import (
	"github.com/daniloAleite/orchestrator/internal/domain/valueobject"
)

type Customer struct {
	ID            string              `json:"id"`
	UserID        string              `json:"user_id"`
	User          *User               `json:"user-example,omitempty"`
	CpfCnpj       valueobject.CpfCnpj `json:"document"`
	PhoneNumber   string              `json:"phone_number"`
	FullName      string              `json:"full_name"`
	Email         string              `json:"email"`
	Vehicles      []Vehicle           `json:"vehicles,omitempty"`
	ServiceOrders []ServiceOrder      `json:"service_orders,omitempty"`
}
