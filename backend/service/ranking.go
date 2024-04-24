/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-13 11:37:56
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 10:15:32
 * @FilePath: /goMall/backend/service/ranking.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"fmt"
	"goMall/backend/pkg/e"
	"goMall/backend/repository/cache"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
	"goMall/backend/serializer"
	"strings"

	logging "github.com/sirupsen/logrus"
)

type RankingService struct {
}

// 获取排行榜
func (service *RankingService) List(c context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.SUCCESS

	// 从redis读取前十
	pros, _ := cache.RedisClient.ZRevRange(cache.RankKey, 0, 9).Result()

	if len(pros) > 1 {
		order := fmt.Sprintf("FIRLD(id, %s)", strings.Join(pros, ","))
		productDao := dao.NewProductDao(c)
		products, err = productDao.SearchForRedis(pros, order)
		if err != nil {
			logging.Info(err)
			code := e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
