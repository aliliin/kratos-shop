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

	CategoryID int32 `gorm:"index:category_id;type:int;comment:分类ID;not null"`
	Category   Category

	Title        string `gorm:"type:varchar(100);comment:属性名;not null"`
	IsAllowAlias bool   `gorm:"default:false;not null;comment:是否允许别名1是0否"`
	IsColor      bool   `gorm:"comment:是否颜色属性;default:false;not null"`
	IsInput      bool   `gorm:"comment:是否输入属性: 1是0否;default:false;not null"`
	IsKey        bool   `gorm:"comment:是否关键属性: 1是0否;default:false;not null"`
	IsSale       bool   `gorm:"comment:是否销售属性:1是0否;default:false;not null"`
	IsSearch     bool   `gorm:"comment:是否搜索字段: 1是0否;default:false;not null"`
	IsMust       bool   `gorm:"comment:是否必须属性: 1是0否;default:false;not null"`
	IsMulti      bool   `gorm:"comment:是否多选: 1是0否;default:false;not null"`
	OnSale       bool   `gorm:"comment:是否上架;default:false;not null"`
	sort         int32  `gorm:"type:int;comment:商品属性排序字段;not null"`
}

// GoodsPropertyValue 商品属性值字表
type GoodsPropertyValue struct {
	BaseFields
	NameID int64 `gorm:"index:property_name_id;type:int;comment:属性名称表ID;not null"`
	Name   GoodsPropertyName
	Title  string `gorm:"type:varchar(100);comment:属性名;not null"`
	Status bool   `gorm:"comment:是否上架;default:false;not null"`
	sort   int32  `gorm:"type:int;comment:商品属性排序字段;not null"`
}
