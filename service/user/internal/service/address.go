package service

import (
	"context"
	v1 "user/api/user/v1"
	"user/internal/biz"
)

func (ua *UserService) CreateAddress(ctx context.Context, req *v1.CreateAddressReq) (*v1.AddressInfo, error) {
	// user is existing
	_, err := ua.uc.UserById(ctx, req.Uid)
	if err != nil {
		return nil, err
	}

	rv, err := ua.ac.AddAddress(ctx, &biz.Address{
		UserID:    req.Uid,
		IsDefault: int(req.IsDefault),
		Mobile:    req.Mobile,
		Name:      req.Name,
		Province:  req.Province,
		City:      req.City,
		Districts: req.Districts,
		Address:   req.Address,
		PostCode:  req.PostCode,
	})

	if err != nil {
		return nil, err
	}

	return &v1.AddressInfo{
		Id:        rv.ID,
		Name:      rv.Name,
		Mobile:    rv.Mobile,
		Province:  rv.Province,
		City:      rv.City,
		Districts: rv.Districts,
		Address:   rv.Address,
		PostCode:  rv.PostCode,
		IsDefault: int32(rv.IsDefault),
	}, nil
}
