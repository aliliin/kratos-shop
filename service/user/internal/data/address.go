package data

import (
	"context"
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
	}, nil
}
