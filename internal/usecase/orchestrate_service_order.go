package usecase

import (
	"context"
	"fmt"
)

type Logger interface {
	Info(msg string, kv ...any)
	Error(msg string, kv ...any)
}

type OSClient interface {
	CreateOS(ctx context.Context, in StartInput) (osID string, err error)
	CancelOS(ctx context.Context, osID string) error
}

type BillingClient interface {
	CreateBudget(ctx context.Context, osID string) (budgetID string, err error)
	CancelBudget(ctx context.Context, budgetID string) error
}

type ExecutionClient interface {
	StartExecution(ctx context.Context, osID string) error
}

type OrchestrateServiceOrder struct {
	log     Logger
	os      OSClient
	billing BillingClient
	exec    ExecutionClient
}

func NewOrchestrateServiceOrder(log Logger, os OSClient, billing BillingClient, exec ExecutionClient) *OrchestrateServiceOrder {
	return &OrchestrateServiceOrder{log: log, os: os, billing: billing, exec: exec}
}

type StartInput struct {
	CustomerID string   `json:"customer_id"`
	VehicleID  string   `json:"vehicle_id"`
	Items      []string `json:"items"`
}

type StartOutput struct {
	OSID     string `json:"os_id"`
	BudgetID string `json:"budget_id"`
	Status   string `json:"status"` // COMPLETED / COMPENSATED / FAILED
}

func (u *OrchestrateServiceOrder) Start(ctx context.Context, in StartInput) (StartOutput, error) {
	u.log.Info("flow started")

	// 1) cria OS
	osID, err := u.os.CreateOS(ctx, in)
	if err != nil {
		u.log.Error("create OS failed", "err", err)
		return StartOutput{Status: "FAILED"}, fmt.Errorf("create OS: %w", err)
	}

	// 2) cria orçamento
	budgetID, err := u.billing.CreateBudget(ctx, osID)
	if err != nil {
		u.log.Error("create budget failed", "os_id", osID, "err", err)
		_ = u.os.CancelOS(ctx, osID) // compensação
		return StartOutput{OSID: osID, Status: "COMPENSATED"}, fmt.Errorf("create budget: %w", err)
	}

	// 3) inicia execução
	if err := u.exec.StartExecution(ctx, osID); err != nil {
		u.log.Error("start execution failed", "os_id", osID, "budget_id", budgetID, "err", err)
		_ = u.billing.CancelBudget(ctx, budgetID)
		_ = u.os.CancelOS(ctx, osID)
		return StartOutput{OSID: osID, BudgetID: budgetID, Status: "COMPENSATED"}, fmt.Errorf("start execution: %w", err)
	}

	u.log.Info("flow completed", "os_id", osID, "budget_id", budgetID)
	return StartOutput{OSID: osID, BudgetID: budgetID, Status: "COMPLETED"}, nil
}
