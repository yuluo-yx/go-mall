package routes

import (
	"github.com/gin-gonic/gin"
	api "go-mall/api/v1"
	"go-mall/middleware"
	"net/http"
)

func NewRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("product/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategories)
		// 轮播图
		v1.GET("carousels", api.ListCarousels)

		// 需要登录保护的操作
		authed := v1.Group("/")
		// 路由中间件
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			// 显示用户金额
			authed.POST("money", api.ShowMoney)

			// 商品相关操作
			authed.POST("product", api.Product)
			authed.GET("product", api.SearchProduct)

			// 收藏夹相关操作
			authed.GET("favorites", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorites)
			authed.DELETE("favorites/:id", api.DeleteFavorites)

			// 地址相关操作
			authed.POST("address", api.CreateAddress)
			authed.GET("address/:id", api.GetAddress)
			authed.GET("address", api.ListAddress)
			authed.PUT("address/:id", api.UpdateAddress)
			authed.DELETE("address/:id", api.DeleteAddress)

			// 购物车相关操作
			authed.POST("carts", api.CreateCarts)
			authed.GET("carts", api.GetCarts)
			authed.PUT("carts/:id", api.UpdateCarts)
			authed.DELETE("carts/:id", api.DeleteCarts)

			// 订单相关操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders/:id", api.ShowOrder)
			authed.GET("orders", api.ListOrders)
			authed.DELETE("orders/:id", api.DeleteOrder)

			// 支付功能
			authed.POST("paydown", api.OrderPay)
		}
	}

	return r
}
