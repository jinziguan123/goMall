/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:46:53
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 21:18:51
 * @FilePath: /goMall/backend/api/v1/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

// @Summary 创建地址
// @Produce json
// @Param name body string true "用户名"
// @Param phone body string false "手机号"
// @Param address body string false "地址"
// @Success 200 {object} string "成功"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/addresses [POST]
func CreateAddress(c *gin.Context) {
	addressService := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Create(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 展示用户最新地址
// @Produce json
// @Param name body string false "用户名"
// @Param phone body string false "手机号"
// @Param address body string false "地址"
// @Success 200 {object} string "成功"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/addresses [GET]
func GetAddress(c *gin.Context) {
	addressService := service.AddressService{}
	res := addressService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(consts.StatusOK, res)
}

// @Summary 展示用户所有收货地址
// @Produce json
// @Param name body string false "用户名"
// @Param phone body string false "手机号"
// @Param address body string false "地址"
// @Success 200 {object} string "成功"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/addresses/{id} [GET]
func ListAddress(c *gin.Context) {
	addressService := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.List(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// 修改收货地址
func UpdateAddress(c *gin.Context) {
	addressService := service.AddressService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// 删除地址
func DeleteAddress(c *gin.Context) {
	addressService := service.AddressService{}
	if err := c.ShouldBind(&addressService); err == nil {
		res := addressService.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
