/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:15:52
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 11:23:29
 * @FilePath: /goMall/backend/repository/database/dao/order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) CreateOrder(order *model.Order) error {
	return dao.DB.Create(*order).Error
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model.Order{}).Where(condition).Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Order("created_at desc").Find(&orders).Error
	return
}

func (dao *OrderDao) GetOrderById(id uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=?", id).First(&order).Error
	return
}

func (dao *OrderDao) DeleteOrderById(id uint) error {
	return dao.DB.Where("id=?", id).Delete(&model.Order{}).Error
}

func (dao *OrderDao) UpdateOrderById(id uint, order *model.Order) error {
	return dao.DB.Where("id=?", id).Updates(order).Error
}
