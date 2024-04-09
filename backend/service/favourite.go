/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-08 11:05:55
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-08 11:17:12
 * @FilePath: /goMall/backend/service/favourite.go
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

type FavoritesService struct {
	ProductId  uint `form:"product_id" json:"product_id"`
	BossId     uint `form:"boss_id" json:"boss_id"`
	FavoriteId uint `form:"favorite_id" json:"favorite_id"`
	PageNum    int  `form:"pageNum"`
	PageSize   int  `form:"pageSize"`
}

func (service *FavoritesService) Show(c context.Context, uId uint) serializer.Response {
	favouritesDao := dao.NewFavouritesDao(c)
	code := e.SUCCESS
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	favourites, total, err := favouritesDao.ListFavouriteByUserId(uId, service.PageSize, service.PageNum)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavourites(c, favourites), uint(total))
}

func (service *FavoritesService) Create(c context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	favouriteDao := dao.NewFavouritesDao(c)
	exist, _ := favouriteDao.FavouriteExistOrNot(service.ProductId, uId)
	// ???
	if exist {
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(c)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(c)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favourite := &model.Favourite{
		UserID:    uId,
		User:      *user,
		ProductID: service.ProductId,
		Product:   *product,
		BossID:    service.BossId,
		Boss:      *boss,
	}

	favouriteDao = dao.NewFavouritesDaoByDB(favouriteDao.DB)
	err = favouriteDao.CreateFavourite(favourite)
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

// Delete 删除收藏夹
func (service *FavoritesService) Delete(ctx context.Context) serializer.Response {
	code := e.SUCCESS

	favoriteDao := dao.NewFavouritesDao(ctx)
	err := favoriteDao.DeleteFavouriteById(service.FavoriteId)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
	}
}
