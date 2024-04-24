/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 17:09:01
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-17 15:39:50
 * @FilePath: /goMall/backend/service/carousel.go
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

type ListCarouselsService struct {
}

type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

func (service *CreateCarouselService) Create(c context.Context, uId uint) serializer.Response {
	carousel := &model.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.SUCCESS
	carouselDao := dao.NewCarouselDao(c)

	err := carouselDao.CreateCarousel(carousel)
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
		Data:   serializer.BuildCarousel(carousel),
	}
}

func (service *ListCarouselsService) List(c context.Context) serializer.Response {
	code := e.SUCCESS
	carouselsCtx := dao.NewCarouselDao(c)
	carousels, err := carouselsCtx.ListCarousels()
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
