package data

// GoodsSku 商品SKU表
type GoodsSku struct {
	BaseFields
	GoodsID    int64 `gorm:"index:goods_id;type:int;comment:商品ID;not null"`
	Goods      Goods
	Title      string `gorm:"type:varchar(100);comment:规格名称;not null"`
	Num        int64  `gorm:"type:int;comment:商品SKU库存;not null"`
	Price      int64  `gorm:"type:int;comment:商品售价;not null"`
	GoodsCode  string `gorm:"type:varchar(100);comment:商品码;not null"`
	BarCode    string `gorm:"type:varchar(100);comment:条码;not null"`
	OnSale     bool   `gorm:"comment:是否上架;default:false;not null"`
	Properties string `gorm:"type:varchar(255);comment:商品属性表ID,以逗号分隔;not null"`
}
