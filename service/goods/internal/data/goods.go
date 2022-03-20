package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
)

// Goods 商品表
type Goods struct {
	BaseFields
	CategoryID int32 `gorm:"index:category_id;type:int;comment:分类ID;not null"`
	BrandsID   int32 `gorm:"index:brand_id;type:int;comment:品牌ID ;not null"`
	TypeID     int64 `gorm:"index:type_id;type:int;comment:商品类型ID ;not null"`

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

func (p *Goods) ToDomain() *domain.Goods {
	return &domain.Goods{
		ID:              p.ID,
		CategoryID:      p.CategoryID,
		BrandsID:        p.BrandsID,
		TypeID:          p.TypeID,
		Name:            p.Name,
		NameAlias:       p.NameAlias,
		GoodsSn:         p.GoodsSn,
		GoodsTags:       p.GoodsTags,
		MarketPrice:     p.MarketPrice,
		GoodsBrief:      p.GoodsBrief,
		GoodsFrontImage: p.GoodsFrontImage,
		GoodsImages:     p.GoodsImages,
		OnSale:          p.OnSale,
		ShipFree:        p.ShipFree,
		ShipID:          p.ShipID,
		IsNew:           p.IsNew,
		IsHot:           p.IsHot,
		ClickNum:        p.ClickNum,
		SoldNum:         p.SoldNum,
		FavNum:          p.FavNum,
	}
}

func (g goodsRepo) CreateGoods(c context.Context, goods *domain.Goods) (*domain.Goods, error) {
	product := &Goods{
		CategoryID:      goods.CategoryID,
		BrandsID:        goods.BrandsID,
		TypeID:          goods.TypeID,
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

	result := g.data.DB(c).Save(product)
	if result.Error != nil {
		return nil, errors.InternalServer("GOODS_CREATE_ERROR", "商品创建失败")
	}
	return product.ToDomain(), nil
}

func (g goodsRepo) GoodsListByIDs(c context.Context, ids ...int64) ([]*domain.Goods, error) {
	var l []*Goods
	if err := g.data.DB(c).Where("id IN (?)", ids).Find(&l).Error; err != nil {
		return nil, errors.NotFound("GOODS_NOT_FOUND", "商品不存在")
	}
	var res []*domain.Goods
	for _, item := range l {
		res = append(res, item.ToDomain())
	}
	return res, nil
}
