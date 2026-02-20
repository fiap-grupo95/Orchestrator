package usecase

import (
	"context"
	"fmt"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
	logs "github.com/daniloAleite/orchestrator/internal/infrastructure/logger"
	"github.com/daniloAleite/orchestrator/internal/usecase/interfaces"
)

type OrchestrateServiceOrder struct {
	os      interfaces.IOSClient
	billing interfaces.IBillingClient
	exec    interfaces.IExecutionClient
}

func NewOrchestrateServiceOrder(os interfaces.IOSClient, billing interfaces.IBillingClient, exec interfaces.IExecutionClient) *OrchestrateServiceOrder {
	return &OrchestrateServiceOrder{os: os, billing: billing, exec: exec}
}

func (u *OrchestrateServiceOrder) Start(ctx context.Context, in request.StartInput) (response.StartOutput, error) {
	logs.Info("flow started")

	// 1) cria OS
	osID, err := u.os.CreateOS(ctx, in)
	if err != nil {
		logs.Error("create OS failed", err)
		return response.StartOutput{Status: "FAILED"}, fmt.Errorf("create OS: %w", err)
	}

	// 2) cria orçamento
	budgetID, err := u.billing.CreateBudget(ctx, osID)
	if err != nil {
		logs.Error("create budget failed", err)
		_ = u.os.CancelOS(ctx, osID) // compensação
		return response.StartOutput{OSID: osID, Status: "COMPENSATED"}, fmt.Errorf("create budget: %w", err)
	}

	// 3) inicia execução
	if err := u.exec.StartExecution(ctx, osID); err != nil {
		logs.Error("start execution failed", err)
		_ = u.billing.CancelBudget(ctx, budgetID)
		_ = u.os.CancelOS(ctx, osID)
		return response.StartOutput{OSID: osID, BudgetID: budgetID, Status: "COMPENSATED"}, fmt.Errorf("start execution: %w", err)
	}

	logs.Info("flow completed")
	return response.StartOutput{OSID: osID, BudgetID: budgetID, Status: "COMPLETED"}, nil
}
