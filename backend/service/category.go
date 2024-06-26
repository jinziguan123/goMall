/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-08 11:05:27
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-13 13:24:03
 * @FilePath: /goMall/backend/service/category.go
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

type ListCategoriesService struct {
}

type CreateCategoryService struct {
	CategoryName string `form:"category_name" json:"category_name"`
}

func (service *CreateCategoryService) Create(c context.Context) serializer.Response {
	category := model.Category{
		CategoryName: service.CategoryName,
	}
	code := e.SUCCESS
	categoryDao := dao.NewCategoryDao(c)
	err := categoryDao.CreateCategory(&category)
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
		Data:   serializer.BuildCategory(&category),
	}
}

func (service *ListCategoriesService) List(ctx context.Context) serializer.Response {
	code := e.SUCCESS
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.ListCategory()
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
		Data:   serializer.BuildCategories(categories),
	}
}
