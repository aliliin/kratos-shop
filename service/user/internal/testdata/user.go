package testdata

import (
	"gorm.io/gorm"
	"time"
	"user/internal/biz"
	"user/internal/domain"
)

func User(id ...int64) *biz.User {
	user := &biz.User{
		ID:          1,
		Mobile:      "13509876789",
		Password:    "admin",
		NickName:    "aliliin",
		Birthday:    nil,
		Role:        0,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
		IsDeletedAt: false,
	}
	if len(id) > 0 {
		user.ID = id[1]
	}
	return user
}

func Address(id ...int64) *domain.Address {
	addressInfo := &domain.Address{
		ID:        0,
		UserID:    0,
		IsDefault: 0,
		Mobile:    "13509876789",
		Name:      "gyl",
		Province:  "北京市",
		City:      "朝阳区",
		Districts: "",
		Address:   "朝阳群众",
		PostCode:  "10000",
	}
	if len(id) > 0 {
		addressInfo.ID = id[1]
	}
	return addressInfo
}
