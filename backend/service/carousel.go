/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 17:09:01
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 17:09:03
 * @FilePath: /goMall/backend/service/carousel.go
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

type ListCarouselsService struct {
}

func (service *ListCarouselsService) List() serializer.Response {
	code := e.SUCCESS
	carouselsCtx := dao.NewCarouselDao(context.Background())
	carousels, err := carouselsCtx.ListAddress()
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
