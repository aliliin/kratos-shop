package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
)

type GoodsInventory struct {
	BaseFields
	SkuID     int64 `gorm:"index:sku_id;type:int;comment:商品SKU_ID;not null"`
	Inventory int64 `gorm:"type:int;comment:商品库存;not null"`
}

type inventoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewInventoryRepo .
func NewInventoryRepo(data *Data, logger log.Logger) biz.InventoryRepo {
	return &inventoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p *GoodsInventory) ToDomain() *domain.Inventory {
	return &domain.Inventory{
		ID:        p.ID,
		SkuID:     p.SkuID,
		Inventory: p.Inventory,
	}
}

func (i inventoryRepo) Create(ctx context.Context, inventory *domain.Inventory) (*domain.Inventory, error) {
	info := GoodsInventory{
		SkuID:     inventory.SkuID,
		Inventory: inventory.Inventory,
	}
	if err := i.data.DB(ctx).Save(&info).Error; err != nil {
		return nil, errors.InternalServer("INENNTORY_SAVE_ERROR", err.Error())
	}
	return info.ToDomain(), nil
}
