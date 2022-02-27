package data

import (
	"gorm.io/gorm"
	"time"
)

// GoodsType 商品类型表
type GoodsType struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;comment:商品类型名称" json:"name"`
	TypeCode  string         `gorm:"type:varchar(50);not null;comment:商品类型编码" json:"type_code"`
	NameAlias string         `gorm:"type:varchar(50);not null;comment:商品类型别名" json:"name_alias"`
	IsVirtual bool           `gorm:"comment:是否是虚拟商品显示;default:false" json:"is_virtual"`
	Desc      string         `gorm:"type:varchar(50);not null;comment:商品类型描述" json:"desc"`
	Sort      int32          `gorm:"comment:类型排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// GoodsBrandType  商品类型表和商品品牌关联表
type GoodsBrandType struct {
	ID      int32 `gorm:"primarykey;type:int" json:"id"`
	BrandID int32 `gorm:"index:brand_id;type:int;comment:商品品牌ID;not null"`
	TypeID  int32 `gorm:"index:type_id;type:int;comment:商品类型ID;not null"`
}

//商品类型 规格分组和规格参数和规格选项有相关的关系。

// Specifications 规格分组表
type Specifications struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(250);not null;comment:规格名称" json:"name"`
	Code      string         `gorm:"type:varchar(50);not null;comment:规格编码" json:"code"`
	TypeID    int32          `gorm:"index:type_id;type:int;comment:商品类型ID;not null"`
	Sort      int32          `gorm:"comment:规格排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// SpecificationsParam 规格参数信息表
type SpecificationsParam struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	SID       int32          `gorm:"index:s_id;type:int;comment:规格组ID;not null"`
	TypeID    int32          `gorm:"index:type_id;type:int;comment:商品类型ID;not null"`
	Name      string         `gorm:"type:varchar(250);not null;comment:规格参数名称" json:"name"`
	Sort      int32          `gorm:"comment:规格排序;default:99;not null;type:int" json:"sort"`
	Status    bool           `gorm:"comment:参数状态;default:false" json:"status"`
	IsSKU     bool           `gorm:"comment:是否通用的SKU持有;default:false" json:"is_sku"`
	IsSelect  bool           `gorm:"comment:是否可查询;default:false" json:"is_select"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// SpecificationsValue 规格参数信息选项表
type SpecificationsValue struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	PID       int32          `gorm:"index:p_id;type:int;comment:规格ID;not null"`
	Value     string         `gorm:"type:varchar(250);not null;comment:规格参数信息值" json:"value"`
	Sort      int32          `gorm:"comment:规格参数值排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
