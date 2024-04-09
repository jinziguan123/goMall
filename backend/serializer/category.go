/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:15:56
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 16:16:09
 * @FilePath: /goMall/backend/serializer/category.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import "goMall/backend/repository/database/model"

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategories(items []*model.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
