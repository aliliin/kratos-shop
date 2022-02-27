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

func (a *addressRepo) DeleteAddress(ctx context.Context, address *biz.Address) error {
	_, err := a.data.uc.DeleteAddress(ctx, &addressService.AddressReq{
		Id:  address.ID,
		Uid: address.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *addressRepo) DefaultAddress(ctx context.Context, address *biz.Address) error {
	_, err := a.data.uc.DefaultAddress(ctx, &addressService.AddressReq{
		Id:  address.ID,
		Uid: address.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *addressRepo) UpdateAddress(c context.Context, address *biz.Address) error {
	_, err := a.data.uc.UpdateAddress(c, &addressService.UpdateAddressReq{
		Id:        address.ID,
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
		return err
	}

	return nil
}

func (a *addressRepo) AddressListByUid(ctx context.Context, uid int64) ([]*biz.Address, error) {
	//var address []Address
	//result := a.data.db.Where(&Address{UserID: uid}).Find(&address)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//
	//if result.RowsAffected == 0 {
	//	return nil, errors.New("地址列表为空")
	//}
	var addressList []*biz.Address
	//for _, v := range address {
	//	addressTmp := biz.Address{
	//		ID:        v.ID,
	//		UserID:    v.UserID,
	//		IsDefault: v.IsDefault,
	//		Mobile:    v.Mobile,
	//		Name:      v.Name,
	//		Province:  v.Province,
	//		City:      v.City,
	//		Districts: v.Districts,
	//		Address:   v.Address,
	//		PostCode:  v.PostCode,
	//	}
	//	addressList = append(addressList, &addressTmp)
	//}
	return addressList, nil
}
