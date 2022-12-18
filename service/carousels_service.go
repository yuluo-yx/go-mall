package service

import (
	"context"
	"go-mall/dao"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
)

type CarouselsService struct {
}

// List 获取轮播图图片
func (service CarouselsService) List(ctx context.Context) serializer.Response {
	code := e.Success

	// 获取dao对象
	carouselDao := dao.NewCarouselsDao(ctx)
	carousels, err := carouselDao.ListCarousels()
	if err != nil {

		//加入日志记录
		util.LogrusObj.Info("err: ", err)

		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
			Data:   "轮播图数据获取失败",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels))),
	}

}
