package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
	"goods/internal/biz"
)

// GoodsSku 商品SKU 表
type GoodsSku struct {
	BaseFields
	GoodsID        int64  `gorm:"index:goods_id;type:int;comment:商品ID;not null"`
	GoodsSn        string `gorm:"type:varchar(100);not null;comment:商品编号"`
	GoodsName      string `gorm:"type:varchar(100);not null;comment:商品名称"`
	Goods          Goods
	SkuName        string `gorm:"type:varchar(100);comment:SKU名称;not null"`
	SkuCode        string `gorm:"type:varchar(100);comment:SKUCode;not null"`
	BarCode        string `gorm:"type:varchar(100);comment:条码;not null"`
	Price          int64  `gorm:"type:int;comment:商品售价;not null"`
	PromotionPrice int64  `gorm:"type:int;comment:商品促销售价;not null"`
	Points         int64  `gorm:"type:int;comment:赠送积分;not null"`
	RemarksInfo    string `gorm:"type:varchar(100);comment:备注信息;not null"`

	Title  string `gorm:"type:varchar(100);comment:规格名称;not null"`
	Pic    string `gorm:"type:varchar(500);not null;comment:规格参数对应的图片" json:"pic"`
	Num    int64  `gorm:"type:int;comment:商品SKU库存;not null"`
	OnSale bool   `gorm:"comment:是否上架;default:false;not null"`
}

type GoodsInventory struct {
	BaseFields
	SkuID     int64 `gorm:"index:sku_id;type:int;comment:商品SKU_ID;not null"`
	SKU       GoodsSku
	Inventory int64 `gorm:"type:int;comment:商品库存;not null"`
	Sale      int64 `gorm:"type:int;comment:商品销量;not null"`
}

// GoodsSpecificationSku 商品规格和商品Sku关联表
type GoodsSpecificationSku struct {
	BaseFields
	SkuID   int64  `gorm:"index:sku_id;type:int;comment:商品SKU_ID;not null"`
	SkuCode string `gorm:"type:varchar(100);comment:商品SKU_Code;not null"`
	Sort    int32  `gorm:"index:sort;type:int;comment:商品参数展示排序;not null"`

	SpecificationId int64 `gorm:"index:specification_id;type:int;comment:商品规格ID;not null"`
	ValueId         int64 `gorm:"index:s_value_id;type:int;comment:商品规格值表ID;not null"`
}

// GoodsAttrSku 商品属性和商品Sku关联表
type GoodsAttrSku struct {
	BaseFields
	SkuID   int64  `gorm:"index:sku_id;type:int;comment:商品SKU_ID;not null"`
	SkuCode string `gorm:"type:varchar(100);comment:商品SKU_Code;not null"`
	Sort    int32  `gorm:"index:sort;type:int;comment:商品参数展示排序;not null"`

	AttrId  int64 `gorm:"index:attr_id;type:int;comment:商品属性ID;not null"`
	ValueId int64 `gorm:"index:attr_value_id;type:int;comment:属性值表ID;not null"`
}

type goodsSkuRepo struct {
	data *Data
	log  *log.Helper
}

// NewGoodsSkuRepoRepo .
func NewGoodsSkuRepoRepo(data *Data, logger log.Logger) biz.GoodsSkuRepo {
	return &goodsSkuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g *goodsSkuRepo) Create(ctx context.Context, req []*biz.Sku) ([]*biz.Sku, error) {

	return nil, nil
}
