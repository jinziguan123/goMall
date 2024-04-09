/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:59:09
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 15:21:12
 * @FilePath: /goMall/backend/routes/route.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routes

import (
	"goMall/backend/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", api.)
	}
}
