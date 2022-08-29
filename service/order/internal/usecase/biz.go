package usecase

import (
	"context"

	"github.com/google/wire"
)

// ProviderSet is usecase providers.
var ProviderSet = wire.NewSet(NewOrderUsecase)

type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}
