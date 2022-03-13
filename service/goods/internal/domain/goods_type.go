package domain

import (
	"strconv"
	"strings"
)

type GoodsType struct {
	ID        int64
	Name      string
	TypeCode  string
	NameAlias string
	IsVirtual bool
	Desc      string
	Sort      int32
	BrandIds  string
}

func (b *GoodsType) IsEmpty() bool {
	return b.BrandIds == ""
}

func (b *GoodsType) FormatBrandIds() ([]int32, error) {
	ids := strings.Replace(b.BrandIds, "，", ",", -1)
	Ids := strings.Split(ids, ",")

	var i []int32
	for _, bid := range Ids {
		if bid == "" {
			continue
		}

		j, err := strconv.ParseInt(bid, 10, 32)
		if err != nil {
			return nil, err
		}
		i = append(i, int32(j))
	}
	return i, nil
}
