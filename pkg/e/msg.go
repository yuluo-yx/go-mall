package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	// 用户错误类型
	ErrorExitUser:                  "用户名已经存在",
	ErrorFailEncryption:            "用户密码加密失败",
	ErrorExitUserNotFound:          "用户不存在",
	ErrorNotCompare:                "用户密码检验错误",
	ErrorAuthToken:                 "token认证失败",
	ErrorAuthCheckTokenTimeout:     "token 已过期",
	ErrorAuthCheckTokenFail:        "token 验证失败",
	ErrorAuthInsufficientAuthority: "token 认证权限不足",
	ErrorUploadFail:                "图片上传失败",

	//邮件错误类型
	ErrorSendEmail: "邮件发送失败",

	// product 商品错误类型
	ErrorProductImgUpload:  "商品图片上传错误",
	ErrorProductCreate:     "保存商品信息失败",
	ErrorProductGet:        "获取商品信息失败",
	ErrorGetProductImgs:    "获取商品图片信息失败",
	ErrorProductCategories: "获取商品分类信息失败",
	ErrorProductExistCart:  "商品不存在",
	ErrorProductMoreCart:   "没有更多的商品",

	// 收藏夹错误类型
	ErrorFavorites:      "获取收藏夹信息失败",
	ErrorFavoritesExist: "收藏夹错误",
	ErrorCreateFavorite: "保存收藏夹信息错误",

	//地址错误类型
	ErrorCreateAddress: "保存地址失败",
	ErrorGetAddress:    "获取地址信息失败",
	ErrorUpdateAddress: "更新地址信息失败",
	ErrorDeleteAddress: "删除地址信息失败",

	//购物车错误类型
	ErrorDeleteCarts: "删除购物车信息失败",
	ErrorUpdateCarts: "更新购物车信息失败",
	ErrorGetCarts:    "获取购物车信息失败",
	ErrorCreateCarts: "保存购物车信息失败",

	// 订单错误类型
	ErrorCreateOrder: "创建订单失败",
	ErrorGetOrder:    "获取订单失败",
	ErrorDeleteOrder: "删除订单失败",

	// 支付错误
	ErrorOrderPayment:      "订单支付失败",
	ErrorLackOfBalance:     "余额不足",
	ErrorUpdateUserBalance: "更新用户余额失败",
	ErrorUpdateBossBalance: "更新老板余额失败",
	ErrorUpdateProductNum:  "更新剩余商品数量失败",
	ErrorCreateUserOrder:   "创建用户已买订单失败",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}

	return msg
}
