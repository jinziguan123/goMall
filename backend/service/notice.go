/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-16 10:20:07
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 10:50:29
 * @FilePath: /goMall/backend/service/notice.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"goMall/backend/pkg/e"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
	"goMall/backend/serializer"

	logging "github.com/sirupsen/logrus"
)

type NoticeService struct {
	Text string `json:"text" form:"text"`
}

// Show 公告详情服务
func (service *NoticeService) Show(c context.Context) serializer.Response {
	code := e.SUCCESS
	noticeDao := dao.NewNoticeDao(c)
	notice, err := noticeDao.GetOneNotice()
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
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildNotice(notice),
	}
}

// Create 创建一条公告
func (service *NoticeService) Create(c context.Context) serializer.Response {
	code := e.SUCCESS
	noticeDao := dao.NewNoticeDao(c)
	notice := &model.Notice{
		Text: service.Text,
	}
	err := noticeDao.CreateNotice(notice)
	if err != nil {
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
