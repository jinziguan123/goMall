/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 15:05:13
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 15:55:22
 * @FilePath: /goMall/backend/database/model/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserID   uint   `gorm:primarykey`
	Username string `gorm:"unique"`
	Email    string
	Password string
	NickName string
	Status   string
	Avator   string `gorm:"size:1000"`
	Money    string
}

const (
	PasswordCost        = 12
	Activate     string = "activate"
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) AvatarURL() string {
	return user.Avator
}
