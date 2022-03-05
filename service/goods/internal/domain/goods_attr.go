package domain

import "goods/internal/biz"

type GoodsAttr struct {
	ID             int64
	TypeID         int32
	GroupID        int64
	Title          string
	Sort           int32
	Status         bool
	Desc           string
	GoodsAttrValue []*biz.GoodsAttrValue
}

type GoodsAttrList []*GoodsAttr

func (p GoodsAttrList) FindById(id int64) *GoodsAttr {
	for _, item := range p {
		if item.ID == id {
			return item
		}
	}
	return nil
}

func (p GoodsAttrList) IsNotExist(groupId, attrId int64) bool {
	for _, item := range p {
		if item.GroupID != groupId && item.ID != attrId {
			return true
		}
	}
	return false
}
