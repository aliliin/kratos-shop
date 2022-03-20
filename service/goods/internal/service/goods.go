package service

import (
	"context"
	v1 "goods/api/goods/v1"
	"goods/internal/domain"
)

// CreateGoods 创建商品
func (g *GoodsService) CreateGoods(ctx context.Context, r *v1.CreateGoodsRequest) (*v1.CreateGoodsResponse, error) {
	var goodsSku []*domain.GoodsSku
	for _, sku := range r.Sku {
		res := &domain.GoodsSku{
			GoodsName:      r.Name,
			GoodsSn:        r.GoodsSn,
			SkuName:        sku.SkuName,
			SkuCode:        sku.Code,
			BarCode:        sku.BarCode,
			Price:          sku.Price,
			PromotionPrice: sku.PromotionPrice,
			Points:         sku.Points,
			Pic:            sku.Image,
			Inventory:      sku.Inventory,
			OnSale:         r.OnSale,
		}

		for _, specification := range sku.SpecificationInfo {
			s := &domain.SpecificationInfo{
				SpecificationID:      specification.SId,
				SpecificationValueID: specification.VId,
			}
			res.Specification = append(res.Specification, s)
		}
		for _, attrGroup := range sku.GroupAttrInfo {
			group := &domain.GroupAttr{
				GroupId:   attrGroup.GroupId,
				GroupName: attrGroup.GroupName,
			}
			for _, attr := range attrGroup.AttrInfo {
				s := &domain.Attr{
					AttrID:        attr.AttrId,
					AttrName:      attr.AttrName,
					AttrValueID:   attr.AttrValueId,
					AttrValueName: attr.AttrValueName,
				}
				group.Attr = append(group.Attr, s)
			}
			res.GroupAttr = append(res.GroupAttr, group)
		}
		goodsSku = append(goodsSku, res)
	}

	goodsInfo := domain.Goods{
		ID:              r.Id,
		CategoryID:      r.CategoryId,
		BrandsID:        r.BrandId,
		TypeID:          r.TypeId,
		Name:            r.Name,
		NameAlias:       r.NameAlias,
		GoodsSn:         r.GoodsSn,
		GoodsTags:       r.GoodsTags,
		MarketPrice:     r.MarketPrice,
		GoodsBrief:      r.GoodsBrief,
		GoodsFrontImage: r.GoodsFrontImage,
		GoodsImages:     r.GoodsImages,
		OnSale:          r.OnSale,
		ShipFree:        r.ShipFree,
		ShipID:          r.ShipId,
		IsNew:           r.IsNew,
		IsHot:           r.IsHot,
		Sku:             goodsSku,
	}

	result, err := g.g.CreateGoods(ctx, &goodsInfo)
	if err != nil {
		return nil, err
	}
	return &v1.CreateGoodsResponse{ID: result.GoodsID}, nil

}

func (g *GoodsService) GoodsList(ctx context.Context, r *v1.GoodsFilterRequest) (*v1.GoodsListResponse, error) {
	goodsFilter := &domain.ESGoodsFilter{
		ID:          r.Id,
		CategoryID:  r.CategoryId,
		BrandsID:    r.BrandId,
		Keywords:    r.Keywords,
		IsNew:       r.IsNew,
		IsHot:       r.IsHot,
		ClickNum:    r.ClickNum,
		SoldNum:     r.SoldNum,
		FavNum:      r.FavNum,
		MaxPrice:    r.MaxPrice,
		MinPrice:    r.MinPrice,
		Pages:       r.Pages,
		PagePerNums: r.PagePerNums,
	}

	result, err := g.esGoods.GoodsList(ctx, goodsFilter)
	if err != nil {
		return nil, err
	}
	response := v1.GoodsListResponse{
		Total: result.Total,
	}
	for _, goods := range result.List {
		res := v1.GoodsInfoResponse{
			Id:          goods.ID,
			CategoryId:  goods.CategoryID,
			BrandId:     goods.BrandsID,
			Name:        goods.Name,
			GoodsSn:     goods.GoodsSn,
			ClickNum:    goods.ClickNum,
			SoldNum:     goods.SoldNum,
			FavNum:      goods.FavNum,
			MarketPrice: goods.MarketPrice,
			GoodsBrief:  goods.GoodsBrief,
			GoodsDesc:   goods.GoodsBrief,
			ShipFree:    goods.ShipFree,
			Images:      goods.GoodsFrontImage,
			GoodsImages: goods.GoodsImages,
			IsNew:       goods.IsNew,
			IsHot:       goods.IsHot,
			OnSale:      goods.OnSale,
		}
		response.List = append(response.List, &res)
	}
	return &response, nil
}
