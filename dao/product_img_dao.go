package dao

import (
	"context"
	"go-mall/model"
	_ "go-mall/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

// NewProductImgDao 创建UserDao 通过上下文的方式创建
func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

// NewProductImgDaoByDB 通过DB去创建ProductDao
func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

// CreateProductImg 创建商品信息
func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

// ListProductImg 获取指定商品id的图片信息 返回一个list
func (dao *ProductImgDao) ListProductImg(id uint) (productImgs []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", id).Find(&productImgs).Error
	return
}
