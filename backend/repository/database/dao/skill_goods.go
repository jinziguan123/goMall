/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:57:13
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 11:57:17
 * @FilePath: /goMall/backend/repository/database/dao/skill_goods.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type SkillGoodsDao struct {
	*gorm.DB
}

func NewSkillGoodsDao(ctx context.Context) *SkillGoodsDao {
	return &SkillGoodsDao{NewDBClient(ctx)}
}

func (dao *SkillGoodsDao) Create(in *model.SkillGoods) error {
	return dao.Model(&model.SkillGoods{}).Create(&in).Error
}

func (dao *SkillGoodsDao) CreateByList(in []*model.SkillGoods) error {
	return dao.Model(&model.SkillGoods{}).Create(&in).Error
}

func (dao *SkillGoodsDao) ListSkillGoods() (resp []*model.SkillGoods, err error) {
	err = dao.Model(&model.SkillGoods{}).Where("num > 0").Find(&resp).Error
	return
}
