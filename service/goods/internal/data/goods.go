package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
)

// Goods 商品表
type Goods struct {
	BaseFields
	CategoryID int32 `gorm:"type:int;comment:分类ID;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;comment:商品ID ;not null"`
	Brands     Brand

	Name            string   `gorm:"type:varchar(100);not null;comment:商品名称"`
	GoodsSn         string   `gorm:"type:varchar(100);not null;comment:商品编号"`
	GoodsTags       string   `gorm:"type:varchar(100);not null;comment:商品标签"`
	MarketPrice     int64    `gorm:"type:int;default:0;not null;comment:商品展示价格"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:商品简介"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:商品封面图"`
	GoodsImages     GormList `gorm:"type:varchar(1000);not null;comment:商品的介绍图"` // 切片类型转为 json 到数据库，取出来是切片类型

	OnSale   bool `gorm:"default:false;comment:是否上架;not null "`
	ShipFree bool `gorm:"default:false;comment:是否免运费; not null"`
	IsNew    bool `gorm:"default:false;comment:是否新品;not null"`
	IsHot    bool `gorm:"comment:是否热卖商品;default:false;not null"`

	ClickNum int64 `gorm:"default:0;type:int; comment 商品详情点击数"`
	SoldNum  int64 `gorm:"default:0;type:int; comment 商品销售数"`
	FavNum   int64 `gorm:"default:0;type:int; comment 商品收藏数"`
}

type goodsRepo struct {
	data *Data
	log  *log.Helper
}

// NewGoodsRepo .
func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *goodsRepo) CreateGreeter(ctx context.Context, g *biz.Goods) error {
	return nil
}

func (r *goodsRepo) UpdateGreeter(ctx context.Context, g *biz.Goods) error {
	return nil
}
