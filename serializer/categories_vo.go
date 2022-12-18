package serializer

import "go-mall/model"

type Categories struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

// BuildCategory 序列化单个Category
func BuildCategory(item *model.Category) Categories {
	return Categories{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

// BuildCategories 序列化整个dCategories
func BuildCategories(items []*model.Category) (categories []Categories) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}

	return categories
}
