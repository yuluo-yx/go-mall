package serializer

import (
	"go-mall/conf"
	"go-mall/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint   `json:"view"`
	CreateAt      int64  `json:"create_at"`
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	BossID        int    `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) Product {
	return Product{
		ID:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          uint(item.View()),
		CreateAt:      item.CreatedAt.Unix(),
		OnSale:        item.OnSale,
		Num:           item.Num,
		BossID:        item.BossID,
		BossName:      item.BossName,
		BossAvatar:    conf.Host + conf.HttpPort + conf.AvatarPath + item.BossAvatar,
	}
}

// BuildProducts 序列化整个Product列表
func BuildProducts(items []*model.Product) (products []Product) {

	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}

	return products
}
