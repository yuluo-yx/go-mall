package e

// 状态码

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// 用户错误类型
	ErrorExitUser                  = 5001
	ErrorFailEncryption            = 5002
	ErrorExitUserNotFound          = 5003
	ErrorNotCompare                = 5004
	ErrorAuthToken                 = 5005
	ErrorAuthCheckTokenTimeout     = 5006
	ErrorAuthCheckTokenFail        = 5007
	ErrorAuthInsufficientAuthority = 5008
	ErrorUploadFail                = 5009

	// 邮件错误类型
	ErrorSendEmail = 6001

	// product 商品错误类型
	ErrorProductImgUpload  = 7001
	ErrorProductCreate     = 7002
	ErrorProductGet        = 7003
	ErrorGetProductImgs    = 7004
	ErrorProductCategories = 7005
	ErrorProductExistCart  = 7006
	ErrorProductMoreCart   = 7007

	// 收藏夹错误类型
	ErrorFavorites      = 8001
	ErrorFavoritesExist = 8002
	ErrorCreateFavorite = 8003

	// 地址错误类型
	ErrorCreateAddress = 9001
	ErrorGetAddress    = 9002
	ErrorUpdateAddress = 9003
	ErrorDeleteAddress = 9004

	// 购物车错误类型
	ErrorDeleteCarts = 10001
	ErrorUpdateCarts = 10002
	ErrorGetCarts    = 10003
	ErrorCreateCarts = 100004

	// 订单错误类型
	ErrorCreateOrder = 11001
	ErrorGetOrder    = 11002
	ErrorDeleteOrder = 11003

	// 支付错误
	ErrorOrderPayment      = 13001
	ErrorLackOfBalance     = 13002
	ErrorUpdateUserBalance = 13003
	ErrorUpdateBossBalance = 13004
	ErrorUpdateProductNum  = 13005
	ErrorCreateUserOrder   = 13006
)
