package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
)

// GoodsAttr 商品属性表
type GoodsAttr struct {
	BaseFields

	GoodsTypeID int32 `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	GoodsType   GoodsType

	Title  string `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc   string `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status bool   `gorm:"comment:状态;default:false;not null"`
	sort   int32  `gorm:"type:int;comment:商品属性排序字段;not null"`
}

type GoodsAttrValue struct {
	BaseFields
	AttrId int64 `gorm:"index:property_name_id;type:int;comment:属性表ID;not null"`
	Attr   GoodsAttr
	value  string `gorm:"type:varchar(100);comment:属性值;not null"`
}

type goodsAttrRepo struct {
	data *Data
	log  *log.Helper
}

// NewGoodsAttrRepo .
func NewGoodsAttrRepo(data *Data, logger log.Logger) biz.GoodsAttrRepo {
	return &goodsAttrRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
