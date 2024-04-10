/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 09:20:56
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 09:21:00
 * @FilePath: /goMall/backend/api/v1/pay.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		utils.LogrusObj.Infoln(err)
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
	}
}
