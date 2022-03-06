package data

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

	Title string `gorm:"type:varchar(100);comment:规格名称;not null"`
	Num   int64  `gorm:"type:int;comment:商品SKU库存;not null"`

	OnSale     bool   `gorm:"comment:是否上架;default:false;not null"`
	Properties string `gorm:"type:varchar(255);comment:商品属性表ID,以逗号分隔;not null"`
}

type GoodsInventory struct {
	BaseFields
	SkuID     int64 `gorm:"index:sku_id;type:int;comment:商品SKU_ID;not null"`
	SKU       GoodsSku
	Inventory int64 `gorm:"type:int;comment:商品库存;not null"`
}
