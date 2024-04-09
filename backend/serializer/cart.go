/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 15:05:19
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 16:15:36
 * @FilePath: /goMall/backend/serializer/cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"context"
	"goMall/backend/config"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
)

type Cart struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check         bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	Desc          string `json:"desc"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	env := config.NewEnv()
	c := Cart{
		ID:            cart.ID,
		UserID:        cart.UserID,
		ProductID:     cart.ProductID,
		CreateAt:      cart.CreatedAt.Unix(),
		Num:           cart.Num,
		MaxNum:        cart.MaxNum,
		Check:         cart.Check,
		Name:          product.Name,
		ImgPath:       env.PhotoHost + env.HttpPort + env.ProductPhotoHost + product.ImgPath,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossName:      boss.UserName,
		Desc:          product.Info,
	}
	return c
}

func BuildCarts(items []*model.Cart) (carts []Cart) {
	for _, item := range items {
		product, err := dao.NewProductDao(context.Background()).GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		boss, err := dao.NewUserDao(context.Background()).GetUserById(item.BossID)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)
	}
	return carts
}
