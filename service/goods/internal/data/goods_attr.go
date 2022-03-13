package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
	"gorm.io/gorm"
	"time"
)

// GoodsAttrGroup  商品属性分组表  手机 -> 主体->屏幕,操作系统,网络支持,基本信息
type GoodsAttrGroup struct {
	ID          int64          `gorm:"primarykey;type:int" json:"id"`
	GoodsTypeID int64          `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	Title       string         `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc        string         `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status      bool           `gorm:"comment:状态;default:false;not null"`
	Sort        int32          `gorm:"type:int;comment:商品属性排序字段;not null"`
	CreatedAt   time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// GoodsAttr 商品属性表 主体->产品名称,上市月份,机身宽度
type GoodsAttr struct {
	ID          int64          `gorm:"primarykey;type:int" json:"id"`
	GoodsTypeID int64          `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	GroupID     int64          `gorm:"index:attr_group_id;type:int;comment:商品属性分组ID;not null"`
	Title       string         `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc        string         `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status      bool           `gorm:"comment:状态;default:false;not null"`
	Sort        int32          `gorm:"type:int;comment:商品属性排序字段;not null"`
	CreatedAt   time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type GoodsAttrValue struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	AttrId    int64          `gorm:"index:property_name_id;type:int;comment:属性表ID;not null"`
	GroupID   int64          `gorm:"index:attr_group_id;type:int;comment:商品属性分组ID;not null"`
	Value     string         `gorm:"type:varchar(100);comment:属性值;not null"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
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

func (p *GoodsAttrGroup) ToDomain() *domain.AttrGroup {
	return &domain.AttrGroup{
		ID:     p.ID,
		TypeID: p.GoodsTypeID,
		Title:  p.Title,
		Desc:   p.Desc,
		Status: p.Status,
		Sort:   p.Sort,
	}
}

func (p *GoodsAttr) ToDomain() *domain.GoodsAttr {
	return &domain.GoodsAttr{
		ID:      p.ID,
		TypeID:  p.GoodsTypeID,
		GroupID: p.GroupID,
		Title:   p.Title,
		Sort:    p.Sort,
		Status:  p.Status,
		Desc:    p.Desc,
	}
}

func (p *GoodsAttrValue) ToDomain() *domain.GoodsAttrValue {
	return &domain.GoodsAttrValue{
		ID:      p.ID,
		AttrId:  p.AttrId,
		GroupID: p.GroupID,
		Value:   p.Value,
	}
}

func (g *goodsAttrRepo) CreateGoodsGroupAttr(ctx context.Context, a *domain.AttrGroup) (*domain.AttrGroup, error) {
	group := GoodsAttrGroup{
		GoodsTypeID: a.TypeID,
		Title:       a.Title,
		Desc:        a.Desc,
		Status:      a.Status,
		Sort:        a.Sort,
	}

	result := g.data.db.Save(&group)
	if result.Error != nil {
		return nil, errors.InternalServer("ATTR_GROUP_SAVE_ERROR", result.Error.Error())
	}

	return group.ToDomain(), nil
}

func (g *goodsAttrRepo) IsExistsGroupByID(ctx context.Context, groupId int64) (*domain.AttrGroup, error) {
	var group GoodsAttrGroup
	if res := g.data.db.First(&group, groupId); res.RowsAffected == 0 {
		return nil, errors.NotFound("ATTR_GROUP_NOT_FOUND", "商品属性组不存在")
	}
	return group.ToDomain(), nil
}

func (g *goodsAttrRepo) CreateGoodsAttr(ctx context.Context, a *domain.GoodsAttr) (*domain.GoodsAttr, error) {
	attr := GoodsAttr{
		GoodsTypeID: a.TypeID,
		GroupID:     a.GroupID,
		Title:       a.Title,
		Desc:        a.Desc,
		Status:      a.Status,
		Sort:        a.Sort,
	}

	if err := g.data.DB(ctx).Save(&attr).Error; err != nil {
		return nil, errors.InternalServer("ATTR_SAVE_ERROR", err.Error())
	}
	return attr.ToDomain(), nil
}

func (g *goodsAttrRepo) CreateGoodsAttrValue(ctx context.Context, r []*domain.GoodsAttrValue) ([]*domain.GoodsAttrValue, error) {
	var attrValue []*GoodsAttrValue
	for _, v := range r {
		attr := GoodsAttrValue{
			AttrId:  v.AttrId,
			GroupID: v.GroupID,
			Value:   v.Value,
		}
		attrValue = append(attrValue, &attr)
	}

	if err := g.data.DB(ctx).Create(&attrValue).Error; err != nil {
		return nil, errors.InternalServer("ATTR_CREATE_ERROR", err.Error())
	}

	var res []*domain.GoodsAttrValue
	for _, v := range attrValue {
		value := v.ToDomain()
		res = append(res, value)
	}
	return res, nil
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.NotFound("ATTR_NOT_FOUND", "商品属性不存在")
		}
		return errors.InternalServer("ATTR_GET_ERROR", result.Error.Error())
	}
	if int64(total) != count {
		return errors.NotFound("ATTR_NOT_FOUND", "部分属性不存在")
	}
	return nil

}

func (g *goodsAttrRepo) ListByIds(ctx context.Context, ids ...int64) (domain.GoodsAttrList, error) {
	var l []*GoodsAttr
	if err := g.data.DB(ctx).Where("id IN (?)", ids).Find(&l).Error; err != nil {
		return nil, errors.NotFound("ATTR_NOT_FOUND", "商品属性不存在")
	}

	var res domain.GoodsAttrList
	for _, item := range l {
		res = append(res, item.ToDomain())
	}
	return res, nil
}
