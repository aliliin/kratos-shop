package data

import (
	"cart/internal/domain"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"time"

	"cart/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ShopCart struct {
	ID         int64          `gorm:"primarykey;type:int" json:"id"`
	UserId     int64          `gorm:"type:int;not null;comment:用户id" json:"user_id"`
	GoodsId    int64          `gorm:"type:int;not null;comment:商品id" json:"goods_id"`
	SkuId      int64          `gorm:"type:int;not null;comment:sku_id" json:"sku_id"`
	GoodsPrice int64          `gorm:"type:int;not null;comment:商品价格" json:"goods_price"`
	GoodsNum   int32          `gorm:"type:int;not null;comment:商品数量" json:"goods_num"`
	GoodsSn    string         `gorm:"type:varchar(500);default:;comment:商品编号"`
	GoodsName  string         `gorm:"type:varchar(500);default:;comment:商品名称"`
	IsSelect   bool           `gorm:"type:tinyint;comment:是否选中;default:false" json:"is_select"`
	CreatedAt  time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *cartRepo) Create(ctx context.Context, c *domain.ShopCart) (*domain.ShopCart, error) {
	cartInfo := ShopCart{
		UserId:     c.UserId,
		GoodsId:    c.GoodsId,
		SkuId:      c.SkuId,
		GoodsPrice: c.GoodsPrice,
		GoodsNum:   c.GoodsNum,
		GoodsSn:    c.GoodsSn,
		GoodsName:  c.GoodsName,
		IsSelect:   c.IsSelect,
	}

	result := r.data.db.Save(&cartInfo)
	if result.Error != nil {
		return nil, errors.NotFound("CREATE_CART_NOT_FOUND", "创建购物车失败")
	}

	return modelToBizResponse(cartInfo), nil
}

func modelToBizResponse(cartInfo ShopCart) *domain.ShopCart {
	return &domain.ShopCart{
		ID:         cartInfo.ID,
		UserId:     cartInfo.UserId,
		GoodsId:    cartInfo.GoodsId,
		SkuId:      cartInfo.SkuId,
		GoodsPrice: cartInfo.GoodsPrice,
		GoodsNum:   cartInfo.GoodsNum,
		GoodsSn:    cartInfo.GoodsSn,
		GoodsName:  cartInfo.GoodsName,
		IsSelect:   cartInfo.IsSelect,
	}
}
