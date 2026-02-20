package interfaces

import (
	"context"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/request"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
)

type IOSClient interface {
	GetOS(ctx context.Context, id string) (*response.ServiceOrderResponse, error)
	CreateOS(ctx context.Context, in request.StartInput) (osID string, err error)
	CancelOS(ctx context.Context, osID string) error
}
