/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-16 10:17:42
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 10:50:45
 * @FilePath: /goMall/backend/api/v1/notice.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// ShowNotice 公告详情
func ShowNotice(c *gin.Context) {
	service := service.NoticeService{}
	res := service.Show(c.Request.Context())
	c.JSON(200, res)
}

// CreateNotice 创建公告
func CreateNotice(c *gin.Context) {
	service := service.NoticeService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// UpdateNotice 更新公告
// func UpdateNotice(c *gin.Context) {
// 	service := service.NoticeService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Update(c.Request.Context())
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }
