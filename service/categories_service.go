package service

import (
	"context"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
)

type CategoriesService struct {
}

func (service *CategoriesService) List(ctx context.Context) serializer.Response {
	code := e.Success
	var err error
	var categories []*model.Category
	var categoriesDao = dao.NewCategoriesDao(ctx)

	// 查询
	categories, err = categoriesDao.ListCategories()
	if err != nil {
		//加入日志记录
		util.LogrusObj.Info("err: ", err)
		code = e.ErrorProductCategories
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
			Data:   "商品分类数据获取失败",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildListResponse(serializer.BuildCategories(categories), uint(len(categories))),
	}
}
