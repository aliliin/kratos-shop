package data

import (
	"context"
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

func (r *cartRepo) Create(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}
