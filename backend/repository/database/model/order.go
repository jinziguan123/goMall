/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 20:05:58
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 20:06:00
 * @FilePath: /goMall/backend/repository/database/model/order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// Order 订单信息
type Order struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	BossID    uint   `gorm:"not null"`
	AddressID uint   `gorm:"not null"`
	Num       int    // 数量
	OrderNum  uint64 // 订单号
	Type      uint   // 1 未支付  2 已支付
	Money     float64
}
