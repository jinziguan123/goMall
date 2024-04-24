/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:57:49
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-11 10:40:04
 * @FilePath: /goMall/backend/repository/database/dao/admin.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type AdminDao struct {
	*gorm.DB
}

func NewAdminDao(c context.Context) *AdminDao {
	return &AdminDao{NewDBClient(c)}
}

func NewAdminDaoByDB(db *gorm.DB) *AdminDao {
	return &AdminDao{db}
}

func (dao *AdminDao) ListAdmins() (admins []*model.Admin, err error) {
	err = dao.DB.Model(&model.User{}).Find(&admins).Error
	return
}

func (dao *AdminDao) GetAdminById(aId uint) (admin *model.Admin, err error) {
	err = dao.DB.Model(&model.Admin{}).Where("id=?", aId).First(&admin).Error
	return
}

func (dao *AdminDao) CreateAdmin(admin *model.Admin) (err error) {
	err = dao.DB.Model(&model.Admin{}).Create(&admin).Error
	return
}

func (dao *AdminDao) DeleteAdminById(aId uint) (err error) {
	err = dao.DB.Model(&model.Admin{}).Where("id=?", aId).Delete(&model.Admin{}).Error
	return
}

func (dao *AdminDao) UpdateAdminById(aId uint, admin *model.Admin) (err error) {
	err = dao.DB.Model(&model.Admin{}).Where("id=?", aId).Updates(&admin).Error
	return
}

func (dao *AdminDao) ExistOrNotByAdminName(adminname string) (admin *model.Admin, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Admin{}).Where("user_name=?", adminname).Count(&count).Error
	if count == 0 {
		return admin, false, err
	}
	err = dao.DB.Model(&model.Admin{}).Where("user_name=?", adminname).First(&admin).Error
	if err != nil {
		return admin, false, err
	}
	return admin, true, err

}
