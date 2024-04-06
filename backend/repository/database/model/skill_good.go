/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 20:06:34
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 20:06:37
 * @FilePath: /goMall/backend/repository/database/model/skill_good.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type SkillGoods struct {
	Id         uint `gorm:"primarykey"`
	ProductId  uint `gorm:"not null"`
	BossId     uint `gorm:"not null"`
	Title      string
	Money      float64
	Num        int `gorm:"not null"`
	CustomId   uint
	CustomName string
}

type SkillGood2MQ struct {
	SkillGoodId uint    `json:"skill_good_id"`
	ProductId   uint    `json:"product_id"`
	BossId      uint    `json:"boss_id"`
	UserId      uint    `json:"user_id"`
	Money       float64 `json:"money"`
	AddressId   uint    `json:"address_id"`
	Key         string  `json:"key"`
}
