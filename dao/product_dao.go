package dao

import (
	"context"
	"go-mall/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

// NewProductDao 创建UserDao 通过上下文的方式创建
func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

// NewProductDaoByDB 通过DB去创建ProductDao
func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// CreateProduct 创建商品信息
func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

// CountProductByCondition 获取该分类下的商品总数
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Count(&total).Error
	return
}

// ListProductsByCondition 获取商品信息
func (dao *ProductDao) ListProductsByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}

// SearchProduct 搜索商品数据
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, count int64, err error) {

	err = dao.DB.Model(&model.Product{}).
		Where("name = ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Count(&count).Error
	if err != nil {
		return
	}

	err = dao.DB.Model(&model.Product{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * (page.PageSize)).
		Limit(page.PageSize).Find(&products).Error

	return
}

// GetProductById 根据商品id获取商品信息
func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", id).First(&product).Error
	return
}

// UpdateProduct 更新商品数量
func (dao *ProductDao) UpdateProduct(pId uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("id = ?", pId).Updates(product).Error
}
