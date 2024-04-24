/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-11 15:16:27
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 00:24:08
 * @FilePath: /goMall/backend/service/product_img.go
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

// CreateImgService 商品图片创建的服务
type ProductImgService struct {
	ProductID uint   `form:"product_id" json:"product_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

func (service *ProductImgService) CreateProductImg(c context.Context, pId uint) serializer.Response {
	code := e.SUCCESS
	productImgDao := dao.NewProductImgDao(c)
	productImg := &model.ProductImg{
		ProductID: service.ProductID,
		ImgPath:   service.ImgPath,
	}
	err := productImgDao.CreateProductImg(productImg)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productImgDao = dao.NewProductImgDaoByDB(productImgDao.DB)
	var productImgs []*model.ProductImg
	productImgs, err = productImgDao.ListProductImgByProductId(pId)
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
		Data:   serializer.BuildProductImgs(productImgs),
		Msg:    e.GetMsg(code),
	}
}

func (service *ProductImgService) List(c context.Context, pId uint) serializer.Response {
	code := e.SUCCESS
	productImgDao := dao.NewProductImgDao(c)
	productImgs, err := productImgDao.ListProductImgByProductId(pId)
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
		Data:   serializer.BuildProductImgs(productImgs),
		Msg:    e.GetMsg(code),
	}
}
