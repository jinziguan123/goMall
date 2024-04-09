/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:05:12
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-08 11:07:42
 * @FilePath: /goMall/backend/repository/database/dao/favourite.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:05:12
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-08 11:07:18
 * @FilePath: /goMall/backend/repository/database/dao/favourite.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type FavoritesDao struct {
	*gorm.DB
}

func NewFavouritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{NewDBClient(ctx)}
}

func NewFavouritesDaoByDB(db *gorm.DB) *FavoritesDao {
	return &FavoritesDao{db}
}

// 通过uId来获取收藏夹列表
func (dao *FavoritesDao) ListFavouriteByUserId(uId uint, pageSize, pageNum int) (favourites []*model.Favourite, total int64, err error) {
	err = dao.DB.Model(&model.Favourite{}).Preload("User").Where("user_id=?", uId).Count(&total).Error
	if err != nil {
		return
	}
	// 分页
	err = dao.DB.Model(&model.Favourite{}).Preload("User").Where("user_id=?", uId).Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Find(&favourites).Error
	return
}

func (dao *FavoritesDao) CreateFavourite(favourite *model.Favourite) (err error) {
	err = dao.DB.Create(&favourite).Error
	return
}

func (dao *FavoritesDao) FavouriteExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favourite{}).Where("product_id=? AND user_id=?", pId, uId).Count(&count).Error
	if count == 0 || err != nil {
		return false, nil
	}
	return true, err
}

func (dao *FavoritesDao) DeleteFavouriteById(fId uint) error {
	return dao.DB.Where("id=?", fId).Delete(&model.Favourite{}).Error
}
