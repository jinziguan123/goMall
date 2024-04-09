/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:47:30
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 17:07:56
 * @FilePath: /goMall/backend/service/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"goMall/backend/pkg/e"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
	"goMall/backend/serializer"
	"strconv"

	logging "github.com/sirupsen/logrus"
)

type AddressService struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

func (service *AddressService) Create(c context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(c)
	address := &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao = dao.NewAddressDaoByDB(addressDao.DB)
	var addresses []*model.Address
	addresses, err = addressDao.ListAddressByUid(uId)
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
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Show(c context.Context, aId string) serializer.Response {
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(c)
	addressId, _ := strconv.Atoi(aId)
	address, err := addressDao.GetAddressByAid(uint(addressId))
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
		Data:   serializer.BuildAddress(address),
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) List(c context.Context, uId uint) serializer.Response {
	code := e.SUCCESS
	addressDao := dao.NewAddressDao(c)
	address, err := addressDao.ListAddressByUid(uId)
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
		Data:   serializer.BuildAddresses(address),
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Delete(c context.Context, aId string) serializer.Response {
	addressDao := dao.NewAddressDao(c)
	code := e.SUCCESS
	addressId, _ := strconv.Atoi(aId)
	err := addressDao.DeleteAddressById(uint(addressId))
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
	}
}

func (service *AddressService) Update(ctx context.Context, uid uint, aid string) serializer.Response {
	code := e.SUCCESS

	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	addressId, _ := strconv.Atoi(aid)
	err := addressDao.UpdateAddressById(uint(addressId), address)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao = dao.NewAddressDaoByDB(addressDao.DB)
	var addresses []*model.Address
	addresses, err = addressDao.ListAddressByUid(uid)
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
		Data:   serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}
}
