package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"gorm.io/gorm"
	"time"
)

// Category 商品分类表
type Category struct {
	ID               int32          `gorm:"primarykey;type:int" json:"id"`
	Name             string         `gorm:"type:varchar(50);not null;comment:分类名称" json:"name"`
	ParentCategoryID int32          `json:"parent_id"`
	ParentCategory   *Category      `json:"-"`
	SubCategory      []*Category    `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32          `gorm:"column:level;default:1;not null;type:int;comment:分类的级别" json:"level"`
	IsTab            bool           `gorm:"comment:是否显示;default:false" json:"is_tab"`
	Sort             int32          `gorm:"comment:分类排序;default:99;not null;type:int" json:"sort"`
	CreatedAt        time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

// GoodsCategoryBrand  商品和分类多对对的表
type GoodsCategoryBrand struct {
	ID int32 `gorm:"primarykey;type:int" json:"id"` // bigint

	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique;comment:商品和分类联合索引唯一"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique:comment:商品和分类联合索引唯一"`
	Brands   Brand

	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CategoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &CategoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CategoryRepo) Category(ctx context.Context) ([]*biz.Categories, error) {
	var cate []*Category
	result := r.data.db.Where(&Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&cate)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("分类为空")
	}
	//var res []*biz.Categories
	//err := copier.Copy(res, &cate)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil

}

//
//// buildData 数据的资源组装
//func buildData(list []*Category) map[int32]map[int32]biz.Categories {
//	var data map[int32]map[int32]biz.Categories = make(map[int32]map[int32]biz.Categories)
//	for _, v := range list {
//		id := v.ID
//		fid := v.ParentCategoryID
//		if _, ok := data[fid]; !ok {
//			data[fid] = make(map[int32]biz.Categories)
//		}
//		data[fid][id] = v.(biz.)
//	}
//	return data
//}
//
//// makeTreeCore 图形化
//func (myL *BusinessRelationLogic) makeTreeCore(index int, data map[int]map[int]models.BusinessRelationOther) []models.BusinessRelationOther {
//	tmp := make([]models.BusinessRelationOther)
//	for id, item := range data[index] {
//		if data[id] != nil {
//			item.List = myL.makeTreeCore(id, data)
//		}
//		tmp = append(tmp, item)
//	}
//	return tmp
//}
