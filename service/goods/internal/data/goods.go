package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
	"goods/internal/biz"
)

// Goods 商品表
type Goods struct {
	BaseFields
	CategoryID int32 `gorm:"type:int;comment:分类ID;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;comment:品牌ID ;not null"`
	Brands     Brand

	TypeID int32 `gorm:"type:int;comment:商品类型ID ;not null"`

	Name            string   `gorm:"type:varchar(100);not null;comment:商品名称"`
	NameAlias       string   `gorm:"type:varchar(100);not null;comment:商品别名"`
	GoodsSn         string   `gorm:"type:varchar(100);not null;comment:商品编号"`
	GoodsTags       string   `gorm:"type:varchar(100);not null;comment:商品标签"`
	MarketPrice     int64    `gorm:"type:int;default:0;not null;comment:商品展示价格"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:商品简介"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:商品封面图"`
	GoodsImages     GormList `gorm:"type:varchar(1000);not null;comment:商品的介绍图"` // 切片类型转为 json 到数据库，取出来是切片类型

	OnSale   bool  `gorm:"default:false;comment:是否上架;not null "`
	ShipFree bool  `gorm:"default:false;comment:是否免运费; not null"`
	ShipID   int32 `gorm:"type:int;comment:运费模版ID;not null"`
	IsNew    bool  `gorm:"default:false;comment:是否新品;not null"`
	IsHot    bool  `gorm:"comment:是否热卖商品;default:false;not null"`

	ClickNum int64 `gorm:"default:0;type:int; comment 商品详情点击数"`
	SoldNum  int64 `gorm:"default:0;type:int; comment 商品销售数"`
	FavNum   int64 `gorm:"default:0;type:int; comment 商品收藏数"`

	// 售前服务信息、售后服务信息、商品促销活动信息
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

func (g goodsRepo) CreateGoods(c context.Context, goods *biz.Goods) (*biz.Goods, error) {
	d := &Goods{
		CategoryID:      goods.CategoryID,
		BrandsID:        goods.BrandsID,
		TypeID:          goods.BrandsID,
		Name:            goods.Name,
		NameAlias:       goods.NameAlias,
		GoodsSn:         goods.GoodsSn,
		GoodsTags:       goods.GoodsTags,
		MarketPrice:     goods.MarketPrice,
		GoodsBrief:      goods.GoodsBrief,
		GoodsFrontImage: goods.GoodsFrontImage,
		GoodsImages:     goods.GoodsImages,
		OnSale:          goods.OnSale,
		ShipFree:        goods.ShipFree,
		ShipID:          goods.ShipID,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
	}

	result := g.data.DB(c).Save(d)
	if result.Error != nil {
		return nil, result.Error
	}
	res := &biz.Goods{
		ID:              d.ID,
		CategoryID:      d.CategoryID,
		BrandsID:        d.BrandsID,
		TypeID:          d.TypeID,
		Name:            d.Name,
		NameAlias:       d.NameAlias,
		GoodsSn:         d.GoodsSn,
		GoodsTags:       d.GoodsTags,
		MarketPrice:     d.MarketPrice,
		GoodsBrief:      d.GoodsBrief,
		GoodsFrontImage: d.GoodsFrontImage,
		GoodsImages:     d.GoodsImages,
		OnSale:          d.OnSale,
		ShipFree:        d.ShipFree,
		ShipID:          d.ShipID,
		IsNew:           d.IsNew,
		IsHot:           d.IsHot,
	}
	return res, nil
}
