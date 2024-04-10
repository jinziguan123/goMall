/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:45:21
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 00:36:29
 * @FilePath: /goMall/backend/serializer/product.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"goMall/backend/config"
	"goMall/backend/consts"
	"goMall/backend/repository/database/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        int    `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

// 序列化商品
func BuildProduct(item *model.Product) Product {
	p := Product{
		ID:            item.ID,
		Name:          item.Name,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       config.PhotoHost + config.HttpPort + config.ProductPhotoPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          item.View(),
		Num:           item.Num,
		OnSale:        item.OnSale,
		CreatedAt:     item.CreatedAt.Unix(),
		BossID:        int(item.BossID),
		BossName:      item.BossName,
		BossAvatar:    config.PhotoHost + config.HttpPort + config.AvatarPath + item.BossAvatar,
	}

	if config.UploadModel == consts.UploadModelOss {
		p.ImgPath = item.ImgPath
		p.BossAvatar = item.BossAvatar
	}

	return p

}

// 序列化商品列表
func BuildProducts(items []*model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
