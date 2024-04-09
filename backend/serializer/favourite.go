/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:17:17
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-08 11:08:02
 * @FilePath: /goMall/backend/serializer/favourite.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"context"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
)

type Favourite struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavourite(favourite *model.Favourite, product *model.Product, user *model.User) Favourite {
	return Favourite{
		UserID:        favourite.UserID,
		ProductID:     favourite.ProductID,
		CreatedAt:     favourite.CreatedAt.Unix(),
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BossID:        user.ID,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

func BuildFavourites(c context.Context, items []*model.Favourite) (favourites []Favourite) {
	productDao := dao.NewProductDao(c)
	userDao := dao.NewUserDao(c)

	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		user, err := userDao.GetUserById(item.UserID)
		if err != nil {
			continue
		}
		favourite := BuildFavourite(item, product, user)
		favourites = append(favourites, favourite)
	}
	return favourites
}
