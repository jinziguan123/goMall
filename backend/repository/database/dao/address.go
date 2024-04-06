/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 21:10:55
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 23:43:40
 * @FilePath: /goMall/backend/repository/database/dao/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/repository/database/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(c context.Context) *AddressDao {
	return &AddressDao{NewDBClient(c)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) GetAddressByAid(aId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&address).Error
	return
}

func (dao *AddressDao) ListAddressByUid(uId uint) (addressList []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", uId).Order("created_at desc").Find(&addressList).Error
	return
}

func (dao *AddressDao) CreateAddress(address *model.Address) (err error) {
	err = dao.DB.Model(&model.Address{}).Create(&address).Error
	return
}

func (dao *AddressDao) DeleteAddressById(aId uint) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).Delete(&model.Address{}).Error
	return
}

func (dao *AddressDao) UpdateAddressById(aId uint, address *model.Address) (err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).Updates(&address).Error
	return
}
