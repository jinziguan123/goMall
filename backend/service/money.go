/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-08 11:17:51
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-08 11:17:58
 * @FilePath: /goMall/backend/service/money.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"goMall/backend/pkg/e"
	"goMall/backend/repository/database/dao"
	"goMall/backend/serializer"

	logging "github.com/sirupsen/logrus"
)

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (service *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}
