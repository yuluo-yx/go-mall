package dao

import (
	"context"
	"go-mall/model"
	_ "go-mall/model"
	"gorm.io/gorm"
)

type FavoritesDao struct {
	*gorm.DB
}

// NewFavoritesDao 创建UserDao 通过上下文的方式创建
func NewFavoritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{NewDBClient(ctx)}
}

// NewFavoritesDaoByDB 通过DB去创建UserDao
func NewFavoritesDaoByDB(db *gorm.DB) *FavoritesDao {
	return &FavoritesDao{db}
}

func (dao *FavoritesDao) ListFavorites(id uint) (favorites []*model.Favorite, err error) {

	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", id).Find(&favorites).Error
	return
}

func (dao *FavoritesDao) FavoriteExistOrNot(pId, id uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id = ? AND user_id = ?", pId, id).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (dao *FavoritesDao) CreateFavorite(favorite *model.Favorite) (err error) {
	return dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
}

func (dao *FavoritesDao) DeleteFavorite(uId, fId uint) (err error) {
	return dao.DB.Model(&model.Favorite{}).
		Where("id = ? AND user_id = ?", fId, uId).
		Delete(&model.Favorite{}).Error
}
