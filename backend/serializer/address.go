/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 14:57:09
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 15:01:37
 * @FilePath: /goMall/backend/serializer/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import "goMall/backend/repository/database/model"

type Address struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Seen     bool   `json:"seen"`
	CreateAt int64  `json:"create_at"`
}

// 创建收货地址
func BuildAddress(item *model.Address) Address {
	return Address{
		ID:       item.ID,
		UserID:   item.UserID,
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		Seen:     false,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildAddresses(items []*model.Address) (addresses []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return addresses
}
