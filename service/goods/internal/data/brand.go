package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"gorm.io/gorm"
	"time"
)

// Brand 商品品牌表
type Brand struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;comment:品牌名称" json:"name"`
	Logo      string         `gorm:"type:varchar(200);default:;comment:品牌Logo图片"`
	Desc      string         `gorm:"type:varchar(500);default:;comment:品牌描述"`
	IsTab     bool           `gorm:"comment:是否显示;default:false" json:"is_tab"`
	Sort      int32          `gorm:"comment:品牌排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type BrandRepo struct {
	data *Data
	log  *log.Helper
}

// NewBrandRepo .
func NewBrandRepo(data *Data, logger log.Logger) biz.BrandRepo {
	return &BrandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *BrandRepo) Create(ctx context.Context, b *biz.Brand) (*biz.Brand, error) {
	brand := &Brand{
		Name:  b.Name,
		Logo:  b.Logo,
		Desc:  b.Desc,
		IsTab: b.IsTab,
		Sort:  b.Sort,
	}
	result := r.data.db.Save(brand)
	res := &biz.Brand{
		ID:    brand.ID,
		Name:  brand.Name,
		Logo:  brand.Logo,
		Desc:  brand.Desc,
		IsTab: brand.IsTab,
		Sort:  brand.Sort,
	}
	return res, result.Error
}

func (r *BrandRepo) GetBradByName(ctx context.Context, name string) (*biz.Brand, error) {
	var brand Brand
	result := r.data.db.Where("name=?", name).First(&brand)
	if result.RowsAffected == 1 {
		return &biz.Brand{
			ID:    brand.ID,
			Name:  brand.Name,
			Logo:  brand.Logo,
			Desc:  brand.Desc,
			IsTab: brand.IsTab,
			Sort:  brand.Sort,
		}, nil
	} else {
		return nil, errors.New("品牌不存在")
	}
}

func (r *BrandRepo) Update(ctx context.Context, b *biz.Brand) error {
	brands := Brand{}
	if result := r.data.db.Where("id=?", b.ID).First(&brands); result.RowsAffected == 0 {
		return errors.New("品牌不存在")
	}

	if b.Name != "" {
		brands.Name = b.Name
	}
	if b.Logo != "" {
		brands.Logo = b.Logo
	}
	if b.IsTab {
		brands.IsTab = b.IsTab
	}
	if b.Sort != 0 {
		brands.Sort = b.Sort
	}
	if b.Desc != "" {
		brands.Desc = b.Desc
	}
	result := r.data.db.Save(&brands)
	return result.Error
}
func (r *BrandRepo) List(ctx context.Context, b *biz.Pagination) ([]*biz.Brand, int64, error) {
	var brands []Brand
	result := r.data.db.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&brands)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var rsp []*biz.Brand
	var total int64
	result = r.data.db.Table("brands").Model(&Brand{}).Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	for _, v := range brands {
		br := &biz.Brand{
			ID:    v.ID,
			Name:  v.Name,
			Logo:  v.Logo,
			IsTab: v.IsTab,
			Sort:  v.Sort,
		}
		rsp = append(rsp, br)
	}
	return rsp, total, nil
}

func (r *BrandRepo) IsBrand(ctx context.Context, ids []int32) error {
	idCount := len(ids)
	if idCount == 0 {
		return errors.New("请选择品牌")
	}
	var count int64
	result := r.data.db.Table("brands").Where("id IN (?)", ids).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if int64(idCount) != count {
		return errors.New("品牌不存在")
	}
	return nil
}

func (r *BrandRepo) IsBrandByID(ctx context.Context, id int32) (*biz.Brand, error) {
	var b Brand
	result := r.data.db.Table("brands").Where("id = ?", id).First(&b)
	if result.Error != nil {
		return nil, result.Error
	}
	res := &biz.Brand{
		ID:    b.ID,
		Name:  b.Name,
		Logo:  b.Logo,
		Desc:  b.Desc,
		IsTab: b.IsTab,
		Sort:  b.Sort,
	}
	return res, nil
}
