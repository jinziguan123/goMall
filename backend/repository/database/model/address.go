/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 15:49:41
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 15:00:05
 * @FilePath: /goMall/backend/database/model/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"type:varchar(50) not null"`
}
