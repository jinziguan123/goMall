/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 18:04:15
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 12:26:49
 * @FilePath: /goMall/backend/service/cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"goMall/backend/pkg/e"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
	"goMall/backend/serializer"
	"strconv"

	logging "github.com/sirupsen/logrus"
)

type CartService struct {
	Id        uint `form:"id" json:"id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	ProductId uint `form:"product_id" json:"product_id"`
	Num       uint `form:"num" json:"num"`
}

func (service *CartService) Create(c context.Context, uId uint) serializer.Response {
	var product *model.Product
	code := e.SUCCESS

	// 判断有没有这个商品
	productDao := dao.NewProductDao(c)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 创建购物车
	cartDao := dao.NewCartDao(c)
	cart, status, _ := cartDao.CreateCart(service.ProductId, uId, service.BossID)
	if status == e.ErrorProductMoreCart {
		return serializer.Response{
			Status: status,
			Msg:    e.GetMsg(status),
		}
	}

	userDao := dao.NewUserDao(c)
	boss, _ := userDao.GetUserById(service.BossID)
	return serializer.Response{
		Status: status,
		Msg:    e.GetMsg(status),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) Show(c context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	cartDao := dao.NewCartDao(c)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildCarts(carts),
		Msg:    e.GetMsg(code),
	}
}

// 更新购物车
func (service *CartService) Update(c context.Context, cId string) serializer.Response {
	code := e.SUCCESS
	cartId, _ := strconv.Atoi(cId)

	cartDao := dao.NewCartDao(c)
	err := cartDao.UpdateCartNumById(uint(cartId), service.Num)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 删除购物车
func (service *CartService) Delete(ctx context.Context, cId string) serializer.Response {
	code := e.SUCCESS
	cartId, _ := strconv.Atoi(cId)

	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartById(uint(cartId))
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
