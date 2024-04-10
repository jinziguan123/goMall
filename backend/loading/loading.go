/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 09:43:32
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 09:43:36
 * @FilePath: /goMall/backend/loading/loading.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package loading

import (
	"goMall/backend/pkg/utils"
	"goMall/backend/repository/cache"
	"goMall/backend/repository/database/dao"
)

func Loading() {
	// es.InitEs() // 如果需要接入ELK可以打开这个注释
	dao.InitMySQL()
	cache.InitCache()
	// mq.InitRabbitMQ() // 如果需要接入RabbitMQ可以打开这个注释
	utils.InitLog()
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
