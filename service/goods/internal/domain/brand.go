package domain

type Brand struct {
	ID    int32
	Name  string
	Logo  string
	Desc  string
	IsTab bool
	Sort  int32
}

func (b *Brand) IsNotExist(id int32) bool {
	return false
}

type BrandList []*Brand

func (b BrandList) FindById(id int32) *Brand {
	for _, item := range b {
		if item.ID == id {
			return item
		}
	}
	return nil
}

func (b BrandList) CheckLength(length int) bool {
	return len(b) == length
}
