package service

import (
	"context"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	var err error

	address = &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err = addressDao.CreateAddress(address)
	if err != nil {
		code = e.ErrorCreateAddress
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

func (service *AddressService) Get(ctx context.Context, aId string) serializer.Response {
	var address *model.Address
	addressId, _ := strconv.Atoi(aId)
	code := e.Success
	var err error
	addressDao := dao.NewAddressDao(ctx)

	address, err = addressDao.GetAddressById(uint(addressId))
	if err != nil {
		code = e.ErrorGetAddress
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
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	var address []*model.Address
	code := e.Success
	var err error
	addressDao := dao.NewAddressDao(ctx)

	address, err = addressDao.ListAddressByUseId(uId)
	if err != nil {
		code = e.ErrorGetAddress
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
		Data:   serializer.BuildListResponse(serializer.BuildAddresses(address), uint(len(address))),
	}
}

func (service *AddressService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	var address *model.Address
	var err error
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)

	address = &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err = addressDao.UpdateAddress(uint(addressId), address)
	if err != nil {
		code = e.ErrorUpdateAddress
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

func (service *AddressService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	var err error
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)

	err = addressDao.DeleteAddressByAddressId(uint(addressId), uId)
	if err != nil {
		code = e.ErrorDeleteAddress
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
