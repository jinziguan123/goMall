/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 09:19:17
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 13:26:07
 * @FilePath: /goMall/backend/api/v1/category.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"goMall/backend/consts"
	"goMall/backend/pkg/utils"
	"goMall/backend/service"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	createCategoryService := service.CreateCategoryService{}
	if err := c.ShouldBind(&createCategoryService); err == nil {
		res := createCategoryService.Create(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func ListCategories(c *gin.Context) {
	listCategoriesService := service.ListCategoriesService{}
	if err := c.ShouldBind(&listCategoriesService); err == nil {
		res := listCategoriesService.List(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
