package service

import (
	"context"
	"encoding/json"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *GoodsService) DeleteCategory(ctx context.Context, r *v1.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := g.cac.DeleteCategory(ctx, &biz.CategoryInfo{
		ID: r.Id,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (g *GoodsService) UpdateCategory(ctx context.Context, r *v1.CategoryInfoRequest) (*emptypb.Empty, error) {
	err := g.cac.UpdateCategory(ctx, &biz.CategoryInfo{
		ID:             r.Id,
		Name:           r.Name,
		ParentCategory: r.ParentCategory,
		Level:          r.Level,
		IsTab:          r.IsTab,
		Sort:           r.Sort,
	})
	return &emptypb.Empty{}, err
}

// CreateCategory 创建分类
func (g *GoodsService) CreateCategory(ctx context.Context, r *v1.CategoryInfoRequest) (*v1.CategoryInfoResponse, error) {
	result, err := g.cac.CreateCategory(ctx, &biz.CategoryInfo{
		Name:           r.Name,
		ParentCategory: r.ParentCategory,
		Level:          r.Level,
		IsTab:          r.IsTab,
		Sort:           r.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CategoryInfoResponse{
		Id:             result.ID,
		Name:           result.Name,
		ParentCategory: result.ParentCategory,
		Level:          result.Level,
		IsTab:          result.IsTab,
		Sort:           result.Sort,
	}, nil
}

func (g *GoodsService) GetAllCategoryList(ctx context.Context, r *emptypb.Empty) (*v1.CategoryListResponse, error) {
	cate, err := g.cac.CategoryList(ctx)
	if err != nil {
		return nil, err
	}
	jsonData, _ := json.Marshal(cate)
	res := &v1.CategoryListResponse{
		JsonData: string(jsonData),
	}
	return res, nil
}

// GetSubCategory 获取子分类
func (g *GoodsService) GetSubCategory(ctx context.Context, r *v1.CategoryListRequest) (*v1.SubCategoryListResponse, error) {
	list, err := g.cac.SubCategoryList(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	categoryListRes := v1.SubCategoryListResponse{}
	categoryListRes.Info = &v1.CategoryInfoResponse{
		Id:             list.Category.ID,
		Name:           list.Category.Name,
		ParentCategory: list.Category.ParentCategory,
		Level:          list.Category.Level,
		IsTab:          list.Category.IsTab,
	}

	var subCategoryResponse []*v1.CategoryInfoResponse
	for _, subC := range list.SubCategory {
		subCategoryResponse = append(subCategoryResponse, &v1.CategoryInfoResponse{
			Id:             subC.ID,
			Name:           subC.Name,
			ParentCategory: subC.ParentCategory,
			Level:          subC.Level,
			IsTab:          subC.IsTab,
		})
	}

	categoryListRes.SubCategory = subCategoryResponse
	return &categoryListRes, nil
}
