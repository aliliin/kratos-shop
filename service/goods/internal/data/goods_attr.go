package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
	"goods/internal/biz"
)

// GoodsAttrGroup  商品属性分组表  手机 -> 主体,屏幕,操作系统,网络支持,基本信息
type GoodsAttrGroup struct {
	BaseFields

	GoodsTypeID int32 `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	GoodsType   GoodsType

	Title  string `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc   string `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status bool   `gorm:"comment:状态;default:false;not null"`
	Sort   int32  `gorm:"type:int;comment:商品属性排序字段;not null"`
}

// GoodsAttr 商品属性表
type GoodsAttr struct {
	BaseFields

	GoodsTypeID int32 `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	GoodsType   GoodsType

	GroupID int64 `gorm:"index:attr_group_id;type:int;comment:商品属性分组ID;not null"`

	Title  string `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc   string `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status bool   `gorm:"comment:状态;default:false;not null"`
	sort   int32  `gorm:"type:int;comment:商品属性排序字段;not null"`
}

type GoodsAttrValue struct {
	BaseFields
	AttrId  int64 `gorm:"index:property_name_id;type:int;comment:属性表ID;not null"`
	Attr    GoodsAttr
	GroupID int64  `gorm:"index:attr_group_id;type:int;comment:商品属性分组ID;not null"`
	value   string `gorm:"type:varchar(100);comment:属性值;not null"`
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

func (g *goodsAttrRepo) CreateGoodsAttr(ctx context.Context, a *biz.GoodsAttr) (*biz.GoodsAttr, error) {
	return nil, nil
}

func (g *goodsAttrRepo) CreateGoodsAttrValue(ctx context.Context, a []*biz.GoodsAttrValue) error {
	return nil
}

func (g *goodsAttrRepo) CreateGoodsGroupAttr(ctx context.Context, a *biz.AttrGroup) (*biz.AttrGroup, error) {
	group := &GoodsAttrGroup{
		GoodsTypeID: a.TypeID,
		Title:       a.Title,
		Desc:        a.Desc,
		Status:      a.Status,
		Sort:        a.Sort,
	}
	var res biz.AttrGroup
	result := g.data.db.Save(group)
	if result.Error != nil {
		return &res, nil
	}
	res = biz.AttrGroup{
		ID:     group.ID,
		TypeID: group.GoodsTypeID,
		Title:  group.Title,
		Desc:   group.Desc,
		Status: group.Status,
		Sort:   group.Sort,
	}
	return &res, nil
}
