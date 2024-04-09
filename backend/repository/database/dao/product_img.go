/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:23:55
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 11:23:58
 * @FilePath: /goMall/backend/repository/database/dao/product_img.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

// CreateProductImg 创建商品图片
func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	err = dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
	return
}

// ListProductImgByProductId 根据商品id获取商品图片
func (dao *ProductImgDao) ListProductImgByProductId(pId uint) (products []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).
		Where("product_id=?", pId).Find(&products).Error
	return
}
