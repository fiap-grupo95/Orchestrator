package interfaces

import (
	"context"
)

type IBillingClient interface {
	CreateBudget(ctx context.Context, osID string) (budgetID string, err error)
	CancelBudget(ctx context.Context, budgetID string) error
}
