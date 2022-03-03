package biz

import (
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
)

type SpecificationInfo struct {
	SpecificationID      int64
	SpecificationValueID int64
}
type AttrInfo struct {
	AttrID      int64
	AttrValueID int64
}
type GoodsSku struct {
	GoodsID        int64
	GoodsSn        string
	GoodsName      string
	SkuName        string
	SkuCode        string
	BarCode        string
	Price          int64
	PromotionPrice int64
	Points         int64
	RemarksInfo    string
	Pic            string
	Num            int64
	OnSale         bool

	Specification []*SpecificationInfo
	Attr          []*AttrInfo
}

type GoodsInfo struct {
	ID              int64
	CategoryID      int32
	BrandsID        int32
	TypeID          int32
	Name            string
	NameAlias       string
	GoodsSn         string
	GoodsTags       string
	MarketPrice     int64
	GoodsBrief      string
	GoodsFrontImage string
	GoodsImages     []string
	OnSale          bool
	ShipFree        bool
	ShipID          int32
	IsNew           bool
	IsHot           bool
	Sku             []*GoodsSku
}

type Goods struct {
	ID              int64
	CategoryID      int32
	BrandsID        int32
	TypeID          int32
	Name            string
	NameAlias       string
	GoodsSn         string
	GoodsTags       string
	MarketPrice     int64
	GoodsBrief      string
	GoodsFrontImage string
	GoodsImages     []string
	OnSale          bool
	ShipFree        bool
	ShipID          int32
	IsNew           bool
	IsHot           bool
}

type GoodsInfoResponse struct {
}

type GoodsRepo interface {
	CreateGoods(ctx context.Context, goods *Goods) (*Goods, error)
}

type GoodsUsecase struct {
	repo    GoodsRepo
	skuRepo GoodsSkuRepo
	tr      Transaction
	cRepo   CategoryRepo
	bRepo   BrandRepo
	tRepo   GoodsTypeRepo
	log     *log.Helper
}

func NewGoodsUsecase(repo GoodsRepo, skuRepo GoodsSkuRepo, tx Transaction, gRepo GoodsTypeRepo, cRepo CategoryRepo, bRepo BrandRepo,
	logger log.Logger) *GoodsUsecase {

	return &GoodsUsecase{
		repo: repo,
		//skuRepo: skuRepo,
		tr:    tx,
		tRepo: gRepo,
		cRepo: cRepo,
		bRepo: bRepo,
		log:   log.NewHelper(logger),
	}
}

func (g GoodsUsecase) CreateGoods(ctx context.Context, r *GoodsInfo) (*GoodsInfoResponse, error) {
	var (
		response *GoodsInfoResponse
	)
	// 判断品牌是否存在
	brandInfo, err := g.bRepo.IsBrandByID(ctx, r.BrandsID)
	if err != nil {
		return nil, errors.New("品牌不存在")
	}

	// 判断分类是否存在
	cInfo, err := g.cRepo.GetCategoryByID(ctx, r.CategoryID)
	if err != nil {
		return nil, errors.New("分类不存在")
	}
	// 判断商品类型是否存在
	typeInfo, err := g.tRepo.GetGoodsTypeByID(ctx, r.TypeID)
	if err != nil {
		return nil, errors.New("商品类型不存在")
	}
	fmt.Println(brandInfo, cInfo, typeInfo)
	fmt.Println(response)
	err = g.tr.ExecTx(ctx, func(ctx context.Context) error {
		// 更新商品表
		goods, err := g.repo.CreateGoods(ctx, &Goods{
			CategoryID:      r.CategoryID,
			BrandsID:        r.BrandsID,
			TypeID:          r.TypeID,
			Name:            r.Name,
			NameAlias:       r.NameAlias,
			GoodsSn:         r.GoodsSn,
			GoodsTags:       r.GoodsTags,
			MarketPrice:     r.MarketPrice,
			GoodsBrief:      r.GoodsBrief,
			GoodsFrontImage: r.GoodsFrontImage,
			GoodsImages:     r.GoodsImages,
			OnSale:          r.OnSale,
			IsNew:           r.IsNew,
			IsHot:           r.IsHot,
			ShipFree:        r.ShipFree,
			ShipID:          r.ShipID,
		})
		if err != nil {
			return err
		}
		var sku []*Sku
		for _, v := range r.Sku {
			res := &Sku{
				GoodsID:        goods.ID,
				GoodsSn:        goods.GoodsSn,
				GoodsName:      goods.Name,
				SkuName:        v.SkuName,
				SkuCode:        "",
				BarCode:        "",
				Price:          0,
				PromotionPrice: 0,
				Points:         0,
				RemarksInfo:    "",
				Pic:            "",
				Num:            0,
				OnSale:         false,
			}
			create, err := g.skuRepo.Create(ctx, res)
			if err != nil {
				return err
			}
			// 更新 sku 关联关系表
			fmt.Println(create)
		}
		// 更新 sku 表
		if err != nil {
			return err
		}
		fmt.Println(sku)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return nil, nil
}
