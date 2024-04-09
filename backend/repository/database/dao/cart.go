/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-06 23:50:31
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-08 11:01:40
 * @FilePath: /goMall/backend/repository/database/dao/cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/pkg/e"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(c context.Context) *CartDao {
	return &CartDao{NewDBClient(c)}
}

func NewCartByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

// 创建cart 商品pId 用户uId 店家bId
func (dao *CartDao) CreateCart(pId, uId, bId uint) (cart *model.Cart, status int, err error) {
	cart, err = dao.GetCartById(pId, uId, bId)
	if err == gorm.ErrRecordNotFound {
		// 没找到，新建一条记录
		cart = &model.Cart{
			UserID:    uId,
			ProductID: pId,
			BossID:    bId,
			Num:       1,
			MaxNum:    10,
			Check:     false,
		}
		err = dao.DB.Create(&cart).Error
		if err != nil {
			return
		}
		return cart, e.SUCCESS, err
	} else if cart.Num < cart.MaxNum {
		// 小于最大Num
		cart.Num++
		err = dao.DB.Save(&cart).Error
		if err != nil {
			return
		}
		return cart, e.ErrorProductExistCart, err
	} else {
		// 大于最大Num
		return cart, e.ErrorProductMoreCart, err
	}

}

func (dao *CartDao) GetCartById(pId, uId, bId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=? AND product_id=? AND boss_id=?", uId, pId, bId).First(&cart).Error
	return
}

func (dao *CartDao) ListCartByUserId(uId uint) (carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uId).Find(&carts).Error
	return
}

func (dao *CartDao) UpdateCartNumById(cId, num uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Update("num", num).Error
}

func (dao *CartDao) DeleteCartById(cId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Delete(&model.Cart{}).Error
}
