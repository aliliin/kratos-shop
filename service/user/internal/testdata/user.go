package testdata

import (
	"user/internal/biz"
)

func User(id ...int64) *biz.User {
	obj := &biz.User{
		ID:     0,
		Mobile: "C",
	}
	if len(id) > 0 {
		obj.ID = id[0]
	}
	return obj
}
