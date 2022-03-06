package data

import (
	"gorm.io/gorm"
	"time"
)

// 商品类型 规格分组和规格参数和规格选项有相关的关系。

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
