/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-16 09:58:33
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 10:16:49
 * @FilePath: /goMall/backend/api/v1/ranking.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

func ListRanking(c *gin.Context) {
	service := service.RankingService{}
	res := service.List(c.Request.Context())
	c.JSON(200, res)
}
