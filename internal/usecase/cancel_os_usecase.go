package usecase

import (
	"context"
	logs "github.com/daniloAleite/orchestrator/internal/infrastructure/logger"
	"github.com/daniloAleite/orchestrator/internal/usecase/interfaces"
)

type ICancelOSUseCase interface {
	CancelServiceOrder(ctx context.Context, ServiceOrderId string) error
}

type CancelOSUseCase struct {
	os      interfaces.IOSClient
	entity  interfaces.IEntityAPIClient
	exec    interfaces.IExecutionClient
	billing interfaces.IBillingClient
}

var _ ICancelOSUseCase = (*CancelOSUseCase)(nil)

func NewCancelOSUseCase(os interfaces.IOSClient, entity interfaces.IEntityAPIClient, exec interfaces.IExecutionClient, billing interfaces.IBillingClient) *CancelOSUseCase {
	return &CancelOSUseCase{os: os, entity: entity, exec: exec, billing: billing}
}

func (u *CancelOSUseCase) CancelServiceOrder(ctx context.Context, ServiceOrderId string) error {
	logger := logs.LoggerWithContext(ctx)

	ServiceOrder, err := u.os.GetOS(ctx, ServiceOrderId)
	if err != nil {
		logger.Error().Err(err).Msg("Error getting service order")
		return err
	}

	err = u.entity.ReleasePartsSupply(ctx, ServiceOrder.PartsSupplies)
	if err != nil {
		logger.Error().Err(err).Msg("Error releasing parts supply")
		return err
	}

	err = u.exec.CancelExecution(ctx, ServiceOrder.ID)
	if err != nil {
		logger.Error().Err(err).Msg("Error canceling execution")
		return err
	}

	err = u.billing.CancelBudget(ctx, ServiceOrder.ID)
	if err != nil {
		logger.Error().Err(err).Msg("Error canceling budget")
		return err
	}

	err = u.os.CancelOS(ctx, ServiceOrder.ID)
	if err != nil {
		logger.Error().Err(err).Msg("Error canceling service order")
		return err
	}

	return nil
}
