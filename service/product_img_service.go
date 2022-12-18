package service

import (
	"context"
	"go-mall/dao"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"strconv"
)

type ProductImgService struct {
}

// List 根据商品id查询商品的图片信息
func (service *ProductImgService) List(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pId, _ := strconv.Atoi(id)
	productImgDao := dao.NewProductImgDao(ctx)

	productImgs, err := productImgDao.ListProductImg(uint(pId))
	if err != nil {
		code = e.ErrorGetProductImgs
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
		Data:   serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs))),
	}
}
