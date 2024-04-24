/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-11 10:26:53
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-11 11:02:57
 * @FilePath: /goMall/backend/serializer/admin.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import "goMall/backend/repository/database/model"

// Admin 用户序列化器
type Admin struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildAdmin 序列化用户
func BuildAdmin(admin *model.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		UserName:  admin.UserName,
		Avatar:    admin.AvatarURL(),
		CreatedAt: admin.CreatedAt.Unix(),
	}
}
