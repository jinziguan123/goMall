/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:46:01
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 16:46:04
 * @FilePath: /goMall/backend/serializer/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"goMall/backend/config"
	"goMall/backend/repository/database/model"
)

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nickname"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	env := config.NewEnv()
	u := &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   env.PhotoHost + env.HttpPort + env.AvatarPath + user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	// if env.UploadModel == consts.UploadModelOss {
	// 	u.Avatar = user.Avatar
	// }

	return u
}

func BuildUsers(items []*model.User) (users []*User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
