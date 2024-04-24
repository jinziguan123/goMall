/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-11 11:06:17
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-17 15:20:40
 * @FilePath: /goMall/backend/api/v1/admin.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

// @Summary 管理员注册
// @Produce json
// @Param user_name body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} string "成功"
// @Failure 10002 {object} string "用户已存在"
// @Failure 10006 {object} string "密码加密失败"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v2/admin/register [POST]
func AdminRegister(c *gin.Context) {
	var adminRegisterService service.AdminService
	if err := c.ShouldBind(&adminRegisterService); err == nil {
		res := adminRegisterService.Register(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 管理员登陆
// @Produce json
// @Param user_name body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} string "成功"
// @Failure 10010 {object} string "用户不存在"
// @Failure 10004 {object} string "密码错误"
// @Failure 30003 {object} string "鉴权失败"
// @Router /api/v2/admin/login [POST]
func AdminLogin(c *gin.Context) {
	var adminLoginService service.AdminService
	if err := c.ShouldBind(&adminLoginService); err == nil {
		res := adminLoginService.Login(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
