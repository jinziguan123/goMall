/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 20:04:27
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 20:04:55
 * @FilePath: /goMall/backend/repository/database/model/carousel.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "github.com/jinzhu/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint `gorm:"not null"`
}
