package data

// GoodsProperty 商品属性表
type GoodsProperty struct {
	BaseFields
	GoodsID int64 `gorm:"index:goods_id;type:int;not null;comment:商品ID"`
	Goods   Goods
	NameID  int64 `gorm:"index:property_name_id;type:int;comment:属性名称表ID;not null"`
	Name    GoodsPropertyName
	ValueID int64 `gorm:"index:prop_value_id;type:int;not null;comment:属性值表ID"`
	Value   GoodsPropertyValue
}

// GoodsPropertyName 商品属性名字表
type GoodsPropertyName struct {
	BaseFields

	GoodsTypeID int32 `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	GoodsType   GoodsType

	Title  string `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc   string `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status bool   `gorm:"comment:状态;default:false;not null"`
	sort   int32  `gorm:"type:int;comment:商品属性排序字段;not null"`
}

type GoodsPropertyValue struct {
	BaseFields
	NameID int64 `gorm:"index:property_name_id;type:int;comment:属性名称表ID;not null"`
	Name   GoodsPropertyName

	value string `gorm:"type:varchar(100);comment:属性值;not null"`
}

// GoodsPropertySku 商品属性和商品Sku关联表
type GoodsPropertySku struct {
	BaseFields

	PropertyID    int64 `gorm:"index:property_id;type:int;comment:商品属性ID;not null"`
	GoodsProperty GoodsProperty

	NameID int64 `gorm:"index:property_name_id;type:int;comment:属性名称表ID;not null"`
	Name   GoodsPropertyName

	SkuID   int64  `gorm:"index:sku_id;type:int;comment:商品SKU_ID;not null"`
	SkuCode string `gorm:"type:varchar(100);comment:商品SKU_Code;not null"`

	Status bool  `gorm:"comment:是否上架;default:false;not null"`
	sort   int32 `gorm:"type:int;comment:商品属性排序字段;not null"`
}
