package serializer

import (
	"go-mall/conf"
	"go-mall/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

// BuildProductImg 序列化单个
func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductID,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

// BuildProductImgs 序列化整个
func BuildProductImgs(item []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range item {
		product := BuildProductImg(item)
		productImgs = append(productImgs, product)
	}

	return productImgs
}
