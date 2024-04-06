/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 20:05:11
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 20:05:13
 * @FilePath: /goMall/backend/repository/database/model/cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "github.com/jinzhu/gorm"

// 购物车模型
type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint `gorm:"not null"`
	BossID    uint
	Num       uint
	MaxNum    uint
	Check     bool
}
