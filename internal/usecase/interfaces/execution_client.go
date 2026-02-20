package interfaces

import (
	"context"
)

type IExecutionClient interface {
	StartExecution(ctx context.Context, osID string) error
	CancelExecution(ctx context.Context, osID string) error
}
