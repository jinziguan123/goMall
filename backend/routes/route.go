/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:59:09
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 10:50:51
 * @FilePath: /goMall/backend/routes/route.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routes

import (
	api "goMall/backend/api/v1"
	"goMall/backend/middleware"
	"net/http"

	_ "goMall/backend/docs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	r.StaticFS("/static", http.Dir("./static"))
	// 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		// 邮箱绑定解绑接口
		v1.POST("user/vaild-email", api.ValidEmail)
		//商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("products/:id", api.ShowProduct)
		//轮播图操作
		v1.GET("carousels", api.ListCarousels)
		//商品图片操作
		v1.GET("imgs/:id", api.ShowProductImgs)
		//商品详情图片操作
		// v1.GET("info-imgs/:id", api.ShowInfoImgs)
		//商品参数图片操作
		// v1.GET("param-imgs/:id", api.ShowParamImgs)
		//分类操作
		v1.GET("categories", api.ListCategories)
		//搜索操作
		v1.POST("searches", api.SearchProducts)
		//排行榜/热门
		v1.GET("rankings", api.ListRanking)
		// v1.GET("elec-rankings", api.ListElecRanking)
		// v1.GET("acce-rankings", api.ListAcceRanking)
		//README操作
		// v1.GET("notices", api.ShowNotice)
		// v1.GET("geetest", api.InitGeetest)
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//验证token
			authed.GET("ping", api.CheckToken)
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("user/sending-email", api.SendEmail)
			// 上传操作
			// authed.POST("avatar", api.UploadToken)
			//收藏夹操作
			authed.GET("favorites/:id", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites", api.DeleteFavorite)
			//订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("user/:id/orders", api.ListOrders)
			authed.GET("orders/:num", api.ShowOrder)
			//购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/:id", api.ShowCarts)
			authed.PUT("carts", api.UpdateCart)
			authed.DELETE("carts", api.DeleteCart)
			//收货地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.ListAddress)
			authed.PUT("addresses", api.UpdateAddress)
			authed.DELETE("addresses", api.DeleteAddress)
			//支付操作
			v1.GET("payments", api.OrderPay)
			//数量操作
			// authed.GET("counts/:id", api.ShowCount)
		}

	}
	v2 := r.Group("/api/v2")
	{
		// 管理员注册
		v2.POST("admin/register", api.AdminRegister)
		// 管理员登录
		v2.POST("admin/login", api.AdminLogin)
		//商品操作
		v2.GET("products", api.ListProducts)
		v2.GET("products/:id", api.ShowProduct)
		//轮播图操作
		v2.GET("carousels", api.ListCarousels)
		//商品图片操作
		v2.GET("imgs/:id", api.ShowProductImgs)
		//分类操作
		v2.GET("categories", api.ListCategories)
		authed2 := v2.Group("/")
		//登录验证
		authed2.Use(middleware.JWTAdmin())
		{
			//商品操作
			authed2.POST("products", api.CreateProduct)
			authed2.DELETE("products/:id", api.DeleteProduct)
			authed2.PUT("products", api.UpdateProduct)
			//轮播图操作
			authed2.POST("carousels", api.CreateCarousel)
			//商品图片操作
			authed2.POST("imgs", api.CreateProductImg)
			//商品详情图片操作
			// authed2.POST("info-imgs", api.CreateInfoImg)
			//商品参数图片操作
			// authed2.POST("param-imgs", api.CreateParamImg)
			//分类操作
			authed2.POST("categories", api.CreateCategory)
			//公告操作
			authed2.POST("notices", api.CreateNotice)
			// authed2.PUT("notices", api.UpdateNotice)
		}
	}
	return r
}
