package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGoodsUsecase,
	NewCategoryUsecase,
	NewBrandUsecase,
	NewGoodsTypeUsecase,
	NewSpecificationUsecase,
)

// Transaction 新增事务接口方法
type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}
