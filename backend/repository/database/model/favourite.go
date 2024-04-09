/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 16:35:04
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-08 11:07:37
 * @FilePath: /goMall/backend/database/model/favourite.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "github.com/jinzhu/gorm"

type Favourite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossID"`
	BossID    uint    `gorm:"not null"`
}
