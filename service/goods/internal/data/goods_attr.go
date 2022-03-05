package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
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
	Sort   int32  `gorm:"type:int;comment:商品属性排序字段;not null"`
}

func (p *GoodsAttr) ToDomain() *domain.GoodsAttr {
	attr := p
	res := &domain.GoodsAttr{
		ID:             attr.ID,
		TypeID:         attr.GoodsTypeID,
		GroupID:        attr.GroupID,
		Title:          attr.Title,
		Sort:           attr.Sort,
		Status:         attr.Status,
		Desc:           attr.Desc,
		GoodsAttrValue: nil,
	}

	return res
}

type GoodsAttrValue struct {
	BaseFields
	AttrId  int64 `gorm:"index:property_name_id;type:int;comment:属性表ID;not null"`
	Attr    GoodsAttr
	GroupID int64  `gorm:"index:attr_group_id;type:int;comment:商品属性分组ID;not null"`
	Value   string `gorm:"type:varchar(100);comment:属性值;not null"`
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

func (g *goodsAttrRepo) CreateGoodsAttr(ctx context.Context, a *domain.GoodsAttr) (*domain.GoodsAttr, error) {
	attr := &GoodsAttr{
		GoodsTypeID: a.TypeID,
		GroupID:     a.GroupID,
		Title:       a.Title,
		Desc:        a.Desc,
		Status:      a.Status,
		Sort:        a.Sort,
	}

	var res domain.GoodsAttr
	result := g.data.DB(ctx).Save(attr)
	if result.Error != nil {
		return &res, nil
	}
	res = domain.GoodsAttr{
		ID:             attr.ID,
		TypeID:         attr.GoodsTypeID,
		GroupID:        attr.GroupID,
		Title:          attr.Title,
		Sort:           attr.Sort,
		Status:         attr.Status,
		Desc:           attr.Desc,
		GoodsAttrValue: nil,
	}

	return &res, nil
}

func (g *goodsAttrRepo) CreateGoodsAttrValue(ctx context.Context, a []*biz.GoodsAttrValue) ([]*biz.GoodsAttrValue, error) {

	var value []*GoodsAttrValue
	for _, v := range a {
		r := &GoodsAttrValue{
			AttrId:  v.AttrId,
			GroupID: v.GroupID,
			Value:   v.Value,
		}
		value = append(value, r)
	}
	var res []*biz.GoodsAttrValue
	result := g.data.DB(ctx).Create(&value)
	if result.Error != nil {
		return res, nil
	}
	for _, v := range value {
		r := &biz.GoodsAttrValue{
			ID:      v.ID,
			AttrId:  v.AttrId,
			GroupID: v.GroupID,
			Value:   v.Value,
		}

		res = append(res, r)
	}
	return res, nil
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

func (g *goodsAttrRepo) GetAttrByIDs(ctx context.Context, ids []*int64) error {
	var attrIDs []*int64
	for _, id := range ids {
		attrIDs = append(attrIDs, id)
	}
	total := len(ids)
	var count int64
	result := g.data.DB(ctx).Where("id IN (?)", attrIDs).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if int64(total) != count {
		return errors.New("部分属性不存在")
	}
	return nil

}

func (g *goodsAttrRepo) ListByIds(ctx context.Context, ids ...*int64) (domain.GoodsAttrList, error) {
	var l []*GoodsAttr
	if err := g.data.DB(ctx).Where("id IN (?)", ids).Take(&l).Error; err != nil {
		return nil, err
	}
	var res domain.GoodsAttrList

	for _, item := range l {
		res = append(res, item.ToDomain())
	}

	return res, nil

}
