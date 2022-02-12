package testdata

import (
	"gorm.io/gorm"
	"time"
	"user/internal/biz"
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
