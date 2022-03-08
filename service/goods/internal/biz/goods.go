package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type GroupAttr struct {
	GroupId int64   `json:"group_id"`
	Attr    []*Attr `json:"attr"`
}
type Attr struct {
	AttrID      int64 `json:"attr_id"`
	AttrValueID int64 `json:"attr_value_id"`
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
	Sku             []*domain.GoodsSku
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
	sRepo   SpecificationRepo
	aRepo   GoodsAttrRepo
	log     *log.Helper
}

func NewGoodsUsecase(repo GoodsRepo, skuRepo GoodsSkuRepo, tx Transaction, gRepo GoodsTypeRepo, cRepo CategoryRepo, bRepo BrandRepo,
	sRepo SpecificationRepo, aRepo GoodsAttrRepo, logger log.Logger) *GoodsUsecase {

	return &GoodsUsecase{
		repo:    repo,
		skuRepo: skuRepo,
		tr:      tx,
		tRepo:   gRepo,
		cRepo:   cRepo,
		bRepo:   bRepo,
		sRepo:   sRepo,
		aRepo:   aRepo,
		log:     log.NewHelper(logger),
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
	typeInfo, err := g.tRepo.GetGoodsTypeByID(ctx, int64(r.TypeID))
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
				SkuCode:        v.SkuCode,
				BarCode:        v.BarCode,
				Price:          v.Price,
				PromotionPrice: v.PromotionPrice,
				Points:         v.Points,
				RemarksInfo:    v.RemarksInfo,
				Pic:            v.Pic,
				Inventory:      v.Num,
				OnSale:         v.OnSale,
			}
			// 验证属性值是否存在
			var AttrID []*int64
			for _, attr := range v.GroupAttr {
				for _, id := range attr.Attr {
					AttrID = append(AttrID, &id.AttrID)
				}
			}
			fmt.Println("AttrID", AttrID)
			err := g.aRepo.GetAttrByIDs(ctx, AttrID)
			if err != nil {
				return err
			}
			// 插入 sku 表
			create, err := g.skuRepo.Create(ctx, res)
			if err != nil {
				return err
			}
			// 更新 sku 规格关联关系表
			var sInfoIDs []*domain.Specification
			var sInfos []*GoodsSpecificationSku
			for _, info := range v.Specification {
				s := &domain.Specification{
					ID: info.SpecificationID,
				}

				si := &GoodsSpecificationSku{
					SkuID:   create.ID,
					SkuCode: create.SkuCode,
					//Sort:            info.Sort,
					SpecificationId: info.SpecificationID,
					ValueId:         info.SpecificationValueID,
				}

				sInfoIDs = append(sInfoIDs, s)
				sInfos = append(sInfos, si)
			}
			// 查询规格是否存在
			err = g.sRepo.GetSpecificationByIDs(ctx, sInfoIDs)
			if err != nil {
				return err
			}
			// 插入商品规格关联关系表
			err = g.skuRepo.CreateSkuRelation(ctx, sInfos)
			if err != nil {
				return err
			}
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
