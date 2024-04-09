/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:24:19
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 11:24:23
 * @FilePath: /goMall/backend/repository/database/dao/product.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	model2 "goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// GetProductById 通过 id 获取product
func (dao *ProductDao) GetProductById(id uint) (product *model2.Product, err error) {
	err = dao.DB.Model(&model2.Product{}).Where("id=?", id).
		First(&product).Error
	return
}

// ListProductByCondition 获取商品列表
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model2.BasePage) (products []*model2.Product, err error) {
	err = dao.DB.Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}

// CreateProduct 创建商品
func (dao *ProductDao) CreateProduct(product *model2.Product) error {
	return dao.DB.Model(&model2.Product{}).Create(&product).Error
}

// CountProductByCondition 根据情况获取商品的数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model2.Product{}).Where(condition).Count(&total).Error
	return
}

// DeleteProduct 删除商品
func (dao *ProductDao) DeleteProduct(pId uint) error {
	return dao.DB.Model(&model2.Product{}).Delete(&model2.Product{}).Error
}

// UpdateProduct 更新商品
func (dao *ProductDao) UpdateProduct(pId uint, product *model2.Product) error {
	return dao.DB.Model(&model2.Product{}).Where("id=?", pId).
		Updates(&product).Error
}

// SearchProduct 搜索商品
func (dao *ProductDao) SearchProduct(info string, page model2.BasePage) (products []*model2.Product, err error) {
	err = dao.DB.Model(&model2.Product{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error
	return
}
