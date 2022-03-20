package biz

import (
	"context"
	"goods/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/olivere/elastic/v7"
)

type EsGoodsRepo interface {
	GoodsList(ctx context.Context, es *domain.EsSearch) ([]int64, int64, error)
	InsertEsGoods(context.Context, domain.ESGoods) error
}

type EsGoodsUsecase struct {
	repo         GoodsRepo
	esRepo       EsGoodsRepo
	categoryRepo CategoryRepo
	log          *log.Helper
}

func NewEsGoodsUsecase(repo GoodsRepo, es EsGoodsRepo, cRepo CategoryRepo, logger log.Logger) *EsGoodsUsecase {
	return &EsGoodsUsecase{
		repo:         repo,
		esRepo:       es,
		categoryRepo: cRepo,
		log:          log.NewHelper(logger),
	}
}

func (g EsGoodsUsecase) GoodsList(ctx context.Context, req *domain.ESGoodsFilter) (*domain.GoodsListResponse, error) {
	// 组织 es 查询条件
	var es domain.EsSearch
	if req.Keywords != "" {
		es.ShouldQuery = append(es.ShouldQuery, elastic.NewMultiMatchQuery(req.Keywords, "name", "goods_brief", "sku.sku_name"))
	}
	if req.IsHot {
		es.ShouldQuery = append(es.ShouldQuery, elastic.NewTermQuery("is_hot", req.IsHot)) // 精确字段查询
	}
	if req.ClickNum > 0 {
		es.ShouldQuery = append(es.ShouldQuery, elastic.NewFieldSort("click_num").Desc()) // 根据某个字段排序
	}
	if req.MinPrice > 0 {
		es.ShouldQuery = append(es.ShouldQuery, elastic.NewRangeQuery("shop_price").Gte(req.MinPrice)) // 区间筛选 gte 大于=
	}
	if req.MaxPrice > 0 {
		es.ShouldQuery = append(es.ShouldQuery, elastic.NewRangeQuery("shop_price").Lte(req.MaxPrice)) // lte 小于=
	}
	if req.BrandsID > 0 {
		es.ShouldQuery = append(es.ShouldQuery, elastic.NewTermQuery("brands_id", req.BrandsID))
	}
	// 通过 category 去查询商品
	if req.CategoryID > 0 {
		// 查询分类是否存在
		cate, err := g.categoryRepo.GetCategoryByID(ctx, req.CategoryID)
		if err != nil {
			return nil, err
		}
		categoryIds, err := g.categoryRepo.GetCategoryAll(ctx, cate.Level, req.CategoryID)
		if err != nil {
			return nil, err
		}
		es.ShouldQuery = append(es.ShouldQuery, elastic.NewTermsQuery("category_id", categoryIds...))
	}
	// 分页处理
	switch {
	case req.PagePerNums > 100:
		req.PagePerNums = 100
	case req.PagePerNums <= 0:
		req.PagePerNums = 10
	}
	if req.Pages == 0 {
		req.Pages = 1
	}
	es.Form = (req.Pages - 1) * req.PagePerNums
	es.Size = req.PagePerNums

	res := &domain.GoodsListResponse{}
	goodsIds, total, err := g.esRepo.GoodsList(ctx, &es)
	if err != nil {
		return nil, err
	}
	res.Total = total
	if err != nil {
		return nil, err
	}
	goodsList, err := g.repo.GoodsListByIDs(ctx, goodsIds...)
	if err != nil {
		return nil, err
	}
	res.List = goodsList
	// TODO 根据返回的商品信息，查询所有分类、查询所有品牌、查询所有sku 的信息进行组合
	return res, nil
}
