package interfaces

import (
	"context"
	"github.com/daniloAleite/orchestrator/internal/adapter/http/dto/response"
)

type IEntityAPIClient interface {
	ReleasePartsSupply(ctx context.Context, partsSupply []response.ServiceOrderPartsSupplyResponse) error
}
