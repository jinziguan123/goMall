/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-06 23:47:20
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-06 23:50:20
 * @FilePath: /goMall/backend/repository/database/dao/carousel.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(c context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(c)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) ListAddress() (carousels []*model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}
