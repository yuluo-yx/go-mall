package service

import (
	"context"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
)

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoriteService) List(ctx context.Context, id uint) serializer.Response {

	code := e.Success
	var err error
	var showFavoritesDao = dao.NewFavoritesDao(ctx)
	var favorites []*model.Favorite

	favorites, err = showFavoritesDao.ListFavorites(id)
	if err != nil {
		util.LogrusObj.Info("err：", err)
		code = e.ErrorFavorites
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(len(favorites)))

}

func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var err error
	var exist bool

	favoriteDao := dao.NewFavoritesDao(ctx)
	exist, err = favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorFavoritesExist
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code := e.ErrorExitUserNotFound
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code := e.ErrorExitUserNotFound
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.ErrorProductGet
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	favorite := &model.Favorite{
		User:      *user,
		UserID:    uId,
		Product:   *product,
		ProductID: service.ProductId,
		Boss:      *boss,
		BossID:    service.BossId,
	}

	// 持久化收藏夹信息
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.ErrorCreateFavorite
		util.LogrusObj.Info("err: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) Delete(ctx context.Context, uId uint, fId uint) serializer.Response {
	code := e.Success
	var deleteFavoritesDao = dao.NewFavoritesDao(ctx)

	err := deleteFavoritesDao.DeleteFavorite(uId, fId)
	if err != nil {
		util.LogrusObj.Info("err：", err)
		code = e.ErrorFavorites
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
