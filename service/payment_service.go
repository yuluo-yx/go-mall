package service

import (
	"context"
	"errors"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"strconv"
)

type OrderPaymentService struct {
	OrderId   uint    `form:"order_id" json:"order_id"`
	Money     float64 `form:"money" json:"money"`
	OrderNo   string  `form:"orderNo" json:"orderNo"`
	ProductID int     `form:"product_id" json:"product_id"`
	PayTime   string  `form:"payTime" json:"payTime" `
	Sign      string  `form:"sign" json:"sign" `
	BossID    int     `form:"boss_id" json:"boss_id"`
	BossName  string  `form:"boss_name" json:"boss_name"`
	Num       int     `form:"num" json:"num"`
	Key       string  `form:"key" json:"key"`
}

// PayDown 支付订单模块
func (service *OrderPaymentService) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)

	// 使用事务来做支付失败时的回滚处理, 资金回退
	tx := orderDao.Begin()

	// 订单
	order, err := orderDao.GetOrderById(service.OrderId, uId)
	if err != nil {
		code = e.ErrorOrderPayment
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	// 用户
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.ErrorExitUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对钱进行解密，减去订单金额，在加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	// 剩余金额不足时
	if moneyFloat-money < 0.0 {
		tx.Rollback()
		code = e.ErrorLackOfBalance
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("余额不足").Error(),
		}
	}
	// 余额足够时
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)
	// 更新用户余额
	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		code = e.ErrorUpdateUserBalance
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("更新用户余额失败").Error(),
		}
	}

	// 老板
	boss, err := userDao.GetUserById(uint(service.BossID))
	if err != nil {
		tx.Rollback()
		code = e.ErrorExitUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 对卖方余额进行操作
	moneyStr = util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	// 类型转换
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)
	// 更新 boss余额
	bossDao := dao.NewUserDaoByDB(userDao.DB)
	err = bossDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		code = e.ErrorUpdateBossBalance
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("更新老板余额失败").Error(),
		}
	}

	// 对应的商品数据量-1
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(service.ProductID))
	if err != nil {
		tx.Rollback()
		code = e.ErrorUpdateProductNum
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("更新剩余商品数量失败").Error(),
		}
	}
	product.Num -= num
	err = productDao.UpdateProduct(uint(service.ProductID), product)
	if err != nil {
		tx.Rollback()
		code = e.ErrorUpdateProductNum
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("更新剩余商品数量失败").Error(),
		}
	}

	// 订单删除
	err = orderDao.DeleteOrderById(service.OrderId, uId)
	if err != nil {
		code = e.ErrorDeleteOrder
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("删除订单失败").Error(),
		}
	}

	productUser := &model.Product{
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		Num:           num,
		OnSale:        false,
		BossID:        int(uId),
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}
	// 买完商品后创建成了自己的商品失败。订单失败，回滚
	err = productDao.CreateProduct(productUser)
	if err != nil {
		tx.Rollback()
		logging.Info(err)
		code = e.ErrorCreateUserOrder
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 提交事务
	tx.Commit()

	// 成功返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
