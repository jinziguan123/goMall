/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-11 15:14:12
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 00:26:41
 * @FilePath: /goMall/backend/api/v1/product_img.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

func CreateProductImg(c *gin.Context) {
	productImgService := service.ProductImgService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&productImgService); err == nil {
		res := productImgService.CreateProductImg(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// 展示所有产品图片
func ShowProductImgs(c *gin.Context) {
	productImgService := service.ProductImgService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&productImgService); err == nil {
		res := productImgService.List(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
