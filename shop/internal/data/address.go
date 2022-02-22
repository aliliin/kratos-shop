package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	addressService "shop/api/service/user/v1"
	"shop/internal/biz"
)

type addressRepo struct {
	data *Data
	log  *log.Helper
}

func NewAddressRepo(data *Data, logger log.Logger) biz.AddressRepo {
	return &addressRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/address")),
	}
}

func (a *addressRepo) CreateAddress(c context.Context, address *biz.Address) (*biz.Address, error) {
	createAddress, err := a.data.uc.CreateAddress(c, &addressService.CreateAddressReq{
		Uid:       address.UserID,
		Name:      address.Name,
		Mobile:    address.Mobile,
		Province:  address.Province,
		City:      address.City,
		Districts: address.Districts,
		Address:   address.Address,
		PostCode:  address.PostCode,
		IsDefault: int32(address.IsDefault),
	})
	if err != nil {
		return nil, err
	}
	res := &biz.Address{
		ID:        createAddress.Id,
		IsDefault: int(createAddress.IsDefault),
		Mobile:    createAddress.Mobile,
		Name:      createAddress.Name,
		Province:  createAddress.Province,
		City:      createAddress.City,
		Districts: createAddress.Districts,
		Address:   createAddress.Address,
		PostCode:  createAddress.PostCode,
	}
	return res, nil
}
