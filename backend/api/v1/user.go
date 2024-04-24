/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 09:33:56
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 22:48:46
 * @FilePath: /goMall/backend/api/v1/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/serializer"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Produce json
// @Param name body string true "用户名"
// @Param nick_name body string true "昵称"
// @Param password body string true "密码"
// @Param key body string true "密钥"
// @Success 200 {object} string "成功"
// @Failure 10002 {object} string "用户已存在"
// @Failure 10006 {object} string "密码加密失败"
// @Failure 40001 {object} string "数据库错误"
// @Failure 500 {object} string "密钥错误"
// @Router /api/v1/user/register [POST]
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 用户登陆
// @Produce json
// @Param name body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} string "成功"
// @Failure 10010 {object} string "用户不存在"
// @Failure 10004 {object} string "密码不正确"
// @Failure 30003 {object} string "Token初始化失败"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/user/login [POST]
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 用户更新信息
// @Produce json
// @Param name body string true "用户名"
// @Param nick_name body string true "昵称"
// @Success 200 {object} string "成功"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/user [PUT]
func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateService); err == nil {
		res := userUpdateService.Update(c.Request.Context(), claims.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	uploadAvatarService := service.UserService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatarService); err == nil {
		res := uploadAvatarService.Post(c.Request.Context(), chaim.ID, file, fileSize)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 发送邮件
// @Produce json
// @Param operationType body string true "操作类型"
// @Param email body string true "对方邮箱"
// @Success 200 {object} string "成功"
// @Failure 10010 {object} string "用户不存在"
// @Failure 30003 {object} string "鉴权失败"
// @Failure 30007 {object} string "发送邮件失败"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/user/sending-email [POST]
func SendEmail(c *gin.Context) {
	var sendEmailService service.SendEmailService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmailService); err == nil {
		res := sendEmailService.Send(c.Request.Context(), chaim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 解绑邮箱
// @Produce json
// @Param Token body string true "Token"
// @Success 200 {object} string "成功"
// @Failure 400 {object} string "非法Token"
// @Failure 30001 {object} string "管理员错误"
// @Failure 30002 {object} string "Token过期"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/user/valid-email [POST]
func ValidEmail(c *gin.Context) {
	var vaildEmailService service.ValidEmailService
	if err := c.ShouldBind(vaildEmailService); err == nil {
		res := vaildEmailService.Valid(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 验证Token
// @Produce json
// @Param Token body string true "Token"
// @Success 200 {object} string "成功"
// @Router /api/v1/ping [GET]
func CheckToken(c *gin.Context) {
	c.JSON(consts.StatusOK, serializer.Response{
		Status: consts.StatusOK,
		Msg:    "ok",
	})
}

// func UploadToken(c *gin.Context) {
// 	var service service.UploadService
// }
