/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 09:16:16
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-17 15:39:40
 * @FilePath: /goMall/backend/api/v1/carousel.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

// @Summary 创建轮播图
// @Produce json
// @Param img_path body string true "轮播图保存路径"
// @Success 200 {object} string "成功"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v2/carousels [POST]
func CreateCarousel(c *gin.Context) {
	createCarouselService := service.CreateCarouselService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCarouselService); err == nil {
		res := createCarouselService.Create(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// @Summary 显示轮播图
// @Produce json
// @Success 200 {object} string "成功"
// @Failure 40001 {object} string "数据库错误"
// @Router /api/v1/carousels [GET]
func ListCarousels(c *gin.Context) {
	listCarouselsService := service.ListCarouselsService{}
	if err := c.ShouldBind(&listCarouselsService); err == nil {
		res := listCarouselsService.List(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
