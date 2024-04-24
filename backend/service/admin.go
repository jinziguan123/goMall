/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-11 10:27:42
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-11 11:06:06
 * @FilePath: /goMall/backend/service/admin.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"goMall/backend/config"
	"goMall/backend/consts"
	"goMall/backend/pkg/e"
	"goMall/backend/pkg/utils"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
	"goMall/backend/serializer"

	logging "github.com/sirupsen/logrus"
)

type AdminService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *AdminService) Register(c context.Context) serializer.Response {
	var admin *model.Admin
	code := e.SUCCESS

	adminDao := dao.NewAdminDao(c)
	_, exist, err := adminDao.ExistOrNotByAdminName(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	admin = &model.Admin{
		UserName: service.UserName,
	}
	if err = admin.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if config.UploadModel == consts.UploadModelOss {
		admin.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	} else {
		admin.Avatar = "avatar.JPG"
	}
	// 创建管理员
	err = adminDao.CreateAdmin(admin)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AdminService) Login(c context.Context) serializer.Response {
	var admin *model.Admin
	code := e.SUCCESS
	adminDao := dao.NewAdminDao(c)
	admin, exist, err := adminDao.ExistOrNotByAdminName(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if !exist {
		logging.Info(err)
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if !admin.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 管理员给予1权限， 普通用户此处为0
	token, err := utils.GenerateToken(admin.ID, service.UserName, 1)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildAdmin(admin), Token: token},
		Msg:    e.GetMsg(code),
	}
}
