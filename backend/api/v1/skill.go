/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 09:25:47
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 09:33:11
 * @FilePath: /goMall/backend/api/v1/skill.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

func ImportSkillGoods(c *gin.Context) {
	var skillGoodsImport service.SkillGoodsImport
	file, _, _ := c.Request.FormFile("file")
	if err := c.ShouldBind(&skillGoodsImport); err == nil {
		res := skillGoodsImport.Import(c.Request.Context(), file)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err, "ImportSkillGoods")
	}
}

func InitSkillGoods(c *gin.Context) {
	var skillGoods service.SkillGoodsService
	if err := c.ShouldBind(&skillGoods); err == nil {
		res := skillGoods.InitSkillGoods(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err, "InitSkillGoods")
	}
}

func SkillGoods(c *gin.Context) {
	var skillGoods service.SkillGoodsService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&skillGoods); err == nil {
		res := skillGoods.SkillGoods(c.Request.Context(), claim.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
