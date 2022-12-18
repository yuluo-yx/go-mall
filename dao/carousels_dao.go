package dao

import (
	"context"
	"go-mall/model"
	"gorm.io/gorm"
)

type CarouselsDao struct {
	*gorm.DB
}

// NewCarouselsDao 创建UserDao 通过上下文的方式创建
func NewCarouselsDao(ctx context.Context) *CarouselsDao {
	return &CarouselsDao{NewDBClient(ctx)}
}

// NewCarouselsDaoByDB 通过DB去创建UserDao
func NewCarouselsDaoByDB(db *gorm.DB) *CarouselsDao {
	return &CarouselsDao{db}
}

// ListCarousels 获取轮播图信息
func (dao *CarouselsDao) ListCarousels() (carousels []*model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}
