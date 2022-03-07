package domain

type GroupAttr struct {
	GroupId int64   `json:"group_id"`
	Attr    []*Attr `json:"attr"`
}
type Attr struct {
	AttrID      int64 `json:"attr_id"`
	AttrValueID int64 `json:"attr_value_id"`
}

type GoodsAttr struct {
	ID             int64
	TypeID         int32
	GroupID        int64
	Title          string
	Sort           int32
	Status         bool
	Desc           string
	GoodsAttrValue []*GoodsAttrValue
}

type GoodsAttrValue struct {
	ID      int64
	AttrId  int64
	GroupID int64
	Value   string
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
