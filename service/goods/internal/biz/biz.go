package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewBrandUsecase,
	NewCategoryUsecase,
	NewGoodsTypeUsecase,
	NewSpecificationUsecase,
	NewGoodsAttrUsecase,
	NewGoodsUsecase,
	NewGoodsSkuUsecase,
	NewInventoryUsecase,
	NewEsGoodsUsecase,
)

// Transaction 新增事务接口方法
type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}
