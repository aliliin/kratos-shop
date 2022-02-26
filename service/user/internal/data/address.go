package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
	"user/internal/biz"
)

type Address struct {
	ID        int64     `gorm:"primarykey"`
	UserID    int64     `gorm:"index:idx_user_id;column:user_id;default:1;type:bigint comment '关联用户ID'"`
	IsDefault int       `gorm:"column:is_default;default:0;type:tinyint comment '是否是默认'"`
	Mobile    string    `gorm:"index:idx_mobile;type:varchar(11) comment '手机号码';not null"`
	Name      string    `gorm:"type:varchar(25) comment '收货用户名称'"`
	Province  string    `gorm:"type:varchar(25) comment '省'"`
	City      string    `gorm:"type:varchar(25) comment '市'"`
	Districts string    `gorm:"type:varchar(25) comment '区县'"`
	Address   string    `gorm:"type:varchar(255) comment '收货详细地址'"`
	PostCode  string    `gorm:"type:varchar(25) comment '邮编'"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
}

func (Address) TableName() string {
	return "user_addresses"
}

type adderessRepo struct {
	data *Data
	log  *log.Helper
}

// NewAddressRepo .
func NewAddressRepo(data *Data, logger log.Logger) biz.AddressRepo {
	return &adderessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *adderessRepo) DeleteAddress(ctx context.Context, r *biz.Address) error {
	var address Address
	result := a.data.db.Where(&Address{ID: r.ID}).First(&address)
	if result.Error != nil {
		return result.Error
	}

	if address.UserID != r.UserID {
		return errors.New("用户ID参数有误")
	}
	//address.DeletedAt = time.Now()

	return a.data.db.Delete(&address).Error
}

func (a *adderessRepo) DefaultAddress(ctx context.Context, r *biz.Address) error {
	var address, addressOld Address
	resCurrDefAdr := a.data.db.Where(&Address{UserID: r.UserID, IsDefault: 1}).First(&addressOld)
	if resCurrDefAdr.Error == nil {
		addressOld.IsDefault = 0
		addressOld.UpdatedAt = time.Now()
		a.data.db.Save(&addressOld)
	}

	result := a.data.db.Where(&Address{ID: r.ID}).First(&address)
	if result.Error != nil {
		return result.Error
	}

	if address.UserID != r.UserID {
		return errors.New("用户ID参数有误")
	}

	address.IsDefault = 1
	address.UpdatedAt = time.Now()

	return a.data.db.Save(&address).Error
}

func (a *adderessRepo) UpdateAddress(ctx context.Context, r *biz.Address) error {
	var address Address
	result := a.data.db.Where(&Address{ID: r.ID}).First(&address)
	if result.Error != nil {
		return result.Error
	}

	if address.UserID != r.UserID {
		return errors.New("用户ID参数有误")
	}

	address.Address = r.Address
	address.City = r.City
	address.Mobile = r.Mobile
	address.Name = r.Name
	address.IsDefault = r.IsDefault
	address.PostCode = r.PostCode
	address.Districts = r.Districts
	address.Province = r.Province
	address.UpdatedAt = time.Now()

	return a.data.db.Save(&address).Error
}

func (a *adderessRepo) AddressListByUid(ctx context.Context, uid int64) ([]*biz.Address, error) {
	var address []Address
	result := a.data.db.Where(&Address{UserID: uid}).Find(&address)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("地址列表为空")
	}
	var addressList []*biz.Address
	for _, v := range address {
		addressTmp := modelToBizResponse(v)
		addressList = append(addressList, addressTmp)
	}
	return addressList, nil
}

func (a *adderessRepo) CreateAddress(c context.Context, r *biz.Address) (*biz.Address, error) {
	addInfo := Address{
		UserID:    r.UserID,
		IsDefault: r.IsDefault,
		Mobile:    r.Mobile,
		Name:      r.Name,
		Province:  r.Province,
		City:      r.City,
		Districts: r.Districts,
		Address:   r.Address,
		PostCode:  r.PostCode,
	}

	result := a.data.db.Save(&addInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return modelToBizResponse(addInfo), nil
}

func modelToBizResponse(addInfo Address) *biz.Address {
	return &biz.Address{
		ID:        addInfo.ID,
		UserID:    addInfo.UserID,
		IsDefault: addInfo.IsDefault,
		Mobile:    addInfo.Mobile,
		Name:      addInfo.Name,
		Province:  addInfo.Province,
		City:      addInfo.City,
		Districts: addInfo.Districts,
		Address:   addInfo.Address,
		PostCode:  addInfo.PostCode,
	}
}
