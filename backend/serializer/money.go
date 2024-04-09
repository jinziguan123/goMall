/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:22:31
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 16:40:08
 * @FilePath: /goMall/backend/serializer/money.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"goMall/backend/pkg/utils"
	"goMall/backend/repository/database/model"
)

type Money struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(item *model.User, key string) Money {
	utils.Encrypt.SetKey(key)
	return Money{
		UserID:    item.ID,
		UserName:  item.UserName,
		UserMoney: utils.Encrypt.AesDecoding(item.Money),
	}
}
