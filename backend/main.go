/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 09:42:25
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 19:36:27
 * @FilePath: /goMall/backend/cmd/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"goMall/backend/config"
	"goMall/backend/loading"
	"goMall/backend/routes"
)

//	@title						goMall电商系统
//	@version					1.0
//	@description				后台API接口文档
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	// Ek1+Ep1==Ek2+Ep2
	config.Init()
	loading.Loading()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}
