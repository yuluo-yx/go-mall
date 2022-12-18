package serializer

import "go-mall/model"

type Carousel struct {
	Id        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
	CreateAt  int64  `json:"create_at"`
}

// BuildCarousel 序列化单个Carousel
func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		Id:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductID,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

// BuildCarousels 序列化整个Carousels
func BuildCarousels(items []*model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}

	return carousels
}
