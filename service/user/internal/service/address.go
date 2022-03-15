package service

import (
	"context"
	v1 "user/api/user/v1"
	"user/internal/domain"
)

func (ua *UserService) DeleteAddress(ctx context.Context, req *v1.AddressReq) (*v1.CheckResponse, error) {
	err := ua.ac.DeleteAddress(ctx, &domain.Address{
		ID:     req.Id,
		UserID: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	return &v1.CheckResponse{Success: true}, nil
}

func (ua *UserService) DefaultAddress(ctx context.Context, req *v1.AddressReq) (*v1.CheckResponse, error) {
	err := ua.ac.DefaultAddress(ctx, &domain.Address{
		ID:     req.Id,
		UserID: req.Uid,
	})

	if err != nil {
		return nil, err
	}

	return &v1.CheckResponse{Success: true}, nil
}

func (ua *UserService) UpdateAddress(ctx context.Context, req *v1.UpdateAddressReq) (*v1.CheckResponse, error) {

	err := ua.ac.UpdateAddress(ctx, &domain.Address{
		ID:        req.Id,
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

	return &v1.CheckResponse{Success: true}, nil
}

func (ua *UserService) ListAddress(ctx context.Context, req *v1.ListAddressReq) (*v1.ListAddressReply, error) {
	rv, err := ua.ac.AddressListByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	res := &v1.ListAddressReply{}
	for _, v := range rv {
		address := v1.AddressInfo{
			Id:        v.ID,
			Name:      v.Name,
			Mobile:    v.Mobile,
			Province:  v.Province,
			City:      v.City,
			Districts: v.Districts,
			Address:   v.Address,
			PostCode:  v.PostCode,
			IsDefault: int32(v.IsDefault),
		}
		res.Results = append(res.Results, &address)
	}
	return res, nil
}
func (ua *UserService) CreateAddress(ctx context.Context, req *v1.CreateAddressReq) (*v1.AddressInfo, error) {
	// user is existing

	_, err := ua.uc.UserById(ctx, req.Uid)
	if err != nil {
		return nil, err
	}

	rv, err := ua.ac.AddAddress(ctx, &domain.Address{
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
