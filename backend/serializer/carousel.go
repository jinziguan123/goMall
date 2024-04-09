/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 15:03:11
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 15:05:03
 * @FilePath: /goMall/backend/serializer/carousel.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import "goMall/backend/repository/database/model"

type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductID uint   `json:"product_id"`
	CreatedAt int64  `json:"created_at"`
}

// 序列化轮播图
func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		ProductID: item.ProductID,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// 序列化轮播图列表
func BuildCarousels(items []*model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
