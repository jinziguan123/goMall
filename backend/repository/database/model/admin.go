/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 15:48:51
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 09:53:50
 * @FilePath: /goMall/backend/database/model/admin.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}

// 设置密码
func (Admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	Admin.PasswordDigest = string(bytes)
	return nil
}

// 校验密码
func (Admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(Admin.PasswordDigest), []byte(password))
	return err == nil
}

// 封面地址
func (Admin *Admin) AvatarURL() string {
	//client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	//bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	//signedGetURL, _ := bucket.SignURL(admin.Avatar, oss.HTTPGet, 24*60*60)
	signedGetURL := "https://github.com/CocaineCong/gin-mall/blob/main/static/imgs/avatar/avatar.JPG"
	return signedGetURL
}
