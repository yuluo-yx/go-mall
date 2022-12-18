package dao

import (
	"context"
	"go-mall/model"
	"gorm.io/gorm"
)

type CategoriesDao struct {
	*gorm.DB
}

// NewCategoriesDao 创建UserDao 通过上下文的方式创建
func NewCategoriesDao(ctx context.Context) *CategoriesDao {
	return &CategoriesDao{NewDBClient(ctx)}
}

// NewCategoriesDaoByDB 通过DB去创建UserDao
func NewCategoriesDaoByDB(db *gorm.DB) *CategoriesDao {
	return &CategoriesDao{db}
}

// ListCategories 获取商品分类信息
func (dao *CategoriesDao) ListCategories() (categories []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&categories).Error
	return
}
